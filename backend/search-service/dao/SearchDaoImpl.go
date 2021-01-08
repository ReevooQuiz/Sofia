package dao

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"search-service/entity"
	"strconv"
)

var (
	mongoUrl string
	mysqlUrl string
)

type SearchDaoImpl struct {
	db      *sql.DB
	session *mgo.Session
}

type TransactionContext struct {
	sqlTx   *sql.Tx
	session *mgo.Session
}

type AnswerActionInfo struct {
	Liked      bool
	Approved   bool
	Approvable bool
}

type SearchUserResult struct {
	Uid      int64  `json:"uid"`
	Banned   bool   `json:"banned"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Profile  string `json:"profile"`
	Icon     string `json:"icon"`
}

type KListItem struct {
	Title string `json:"title"`
	Attr  struct {
		Name   string `json:"name"`
		Value  string `json:"value"`
		Origin string `json:"origin"`
	} `json:"attr"`
}

const (
	PageSize       = 5
	questionFields = "qid,closed,raiser,category,accepted_answer,answer_count,view_count,favorite_count,time,scanned"
	answerFields   = "aid,answerer,qid,view_count,comment_count,criticism_count,like_count,approval_count,time,scanned"
	HotListSize    = 10
	KListSize = 3
	KCardSize = 3
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	_ = godotenv.Load(os.Getenv("WORK_DIR") + "credentials.env")
	mongoUrl = os.Getenv("MONGO_URL")
	mysqlUrl = os.Getenv("MYSQL_URL")
}

func (s *SearchDaoImpl) Begin(read bool) (ctx TransactionContext, err error) {
	var tx *sql.Tx
	if read {
		tx, err = s.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	} else {
		tx, err = s.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	}
	if err != nil {
		return
	}
	ss := s.session.New()
	if ss == nil {
		e := tx.Rollback()
		if e != nil {
			log.Warn(e)
		}
		return ctx, errors.New("failed to create mongo session")
	}
	return TransactionContext{tx, ss}, nil
}

func (s *SearchDaoImpl) Commit(t *TransactionContext) {
	t.session.Close()
	e := t.sqlTx.Commit()
	if e != nil {
		log.Warn(e)
	}
}

func (s *SearchDaoImpl) Rollback(t *TransactionContext) {
	t.session.Close()
	e := t.sqlTx.Rollback()
	if e != nil {
		log.Warn(e)
	}
}

func (s *SearchDaoImpl) Init() (err error) {
	s.db, err = sql.Open("mysql", mysqlUrl)
	if err != nil {
		return err
	}
	s.session, err = mgo.Dial(mongoUrl)
	return err
}

func (s *SearchDaoImpl) Destruct() {
	_ = s.db.Close()
}

func (s *SearchDaoImpl) ParseQuestions(rows *sql.Rows) (result []entity.Questions, err error) {
	var it entity.Questions
	for rows.Next() {
		err = rows.Scan(
			&it.Qid,
			&it.Closed,
			&it.Raiser,
			&it.Category,
			&it.AcceptedAnswer,
			&it.AnswerCount,
			&it.ViewCount,
			&it.FavoriteCount,
			&it.Time)
		if err != nil {
			return
		}
		result = append(result, it)
	}
	return result, nil
}

func (s *SearchDaoImpl) GetBannedWords(ctx TransactionContext) (words []string, err error) {
	var rows *sql.Rows
	rows, err = ctx.sqlTx.Query("select word from ban_words")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var it string
		err = rows.Scan(&it)
		if err != nil {
			return
		}
		words = append(words, it)
	}
	return words, nil
}

func (s *SearchDaoImpl) FindQuestionDetails(ctx TransactionContext, questions []entity.Questions) (questionDetails []entity.QuestionDetails) {
	var findErr error
	var current entity.QuestionDetails
	for _, v := range questions {
		findErr = ctx.session.DB("sofia").C("question_details").FindId(v.Qid).One(&current)
		if findErr != nil {
			log.Warn(findErr)
			questionDetails = append(questionDetails, current)
		} else {
			questionDetails = append(questionDetails, entity.QuestionDetails{})
		}
	}
	return
}

func (s *SearchDaoImpl) AssignLabels(ctx TransactionContext, questions []entity.Questions) (err error) {
	var rows *sql.Rows
	for _, v := range questions {
		rows, err = ctx.sqlTx.Query("select title from labels natural join(select lid from question_labels where qid=?)as L", v.Qid)
		if err != nil {
			return
		}
		v.Labels = make([]string, 0)
		var current string
		for rows.Next() {
			err = rows.Scan(&current)
			if err != nil {
				_ = rows.Close()
				return
			}
			v.Labels = append(v.Labels, current)
		}
		_ = rows.Close()
	}
	return nil
}

func (s *SearchDaoImpl) SearchQuestions(ctx TransactionContext, page int64, text string) (questions []entity.Questions, err error) {
	var rows *sql.Rows
	rows, err = ctx.sqlTx.Query("select "+questionFields+" from questions where match(title)against(? in boolean mode)limit ?, ?",
		text, PageSize*page, PageSize)
	if err != nil {
		return
	}
	defer rows.Close()
	questions, err = s.ParseQuestions(rows)
	if err != nil {
		return
	}
	err = s.AssignLabels(ctx, questions)
	return questions, err
}

func (s *SearchDaoImpl) SearchAnswers(ctx TransactionContext, page int64, text string) (details []entity.AnswerDetails, err error) {
	err = ctx.session.DB("sofia").C("answer_details").
		Find(bson.M{"$text": bson.M{"$search": text}}).
		Skip(int(page * PageSize)).
		Limit(PageSize).
		All(&details)
	return
}

func (s *SearchDaoImpl) ParseAnswers(rows *sql.Rows) (result []entity.Answers, err error) {
	var it entity.Answers
	for rows.Next() {
		err = rows.Scan(
			&it.Aid,
			&it.Answerer,
			&it.Qid,
			&it.CommentCount,
			&it.CriticismCount,
			&it.LikeCount,
			&it.ApprovalCount,
			&it.Time)
		if err != nil {
			return
		}
		result = append(result, it)
	}
	return result, nil
}

func (s *SearchDaoImpl) FindAnswerSkeletons(ctx TransactionContext, details []entity.AnswerDetails) (answers []entity.Answers) {
	res := make([]entity.Answers, len(details))
	for i, v := range details {
		var rows *sql.Rows
		rows, err := ctx.sqlTx.Query("select "+answerFields+" from answers where aid=?", v.Aid)
		suc := false
		if err == nil {
			items, err := s.ParseAnswers(rows)
			if err == nil {
				if len(items) > 0 {
					suc = true
					res[i] = items[0]
				} else {
					log.Warn("aid = ", v.Aid, ", answer skeleton not found")
				}
			} else {
				log.Warn(err)
			}
			_ = rows.Close()
		} else {
			log.Warn(err)
		}
		if !suc {
			res[i].Aid = v.Aid
		}
	}
	return res
}

func (s *SearchDaoImpl) GetAnswerActionInfos(ctx TransactionContext, uid int64, qids []int64, aids []int64) (infos []AnswerActionInfo, err error) {
	infos = make([]AnswerActionInfo, len(aids))
	var userLabels = make(map[int64]bool)
	var rows *sql.Rows
	rows, err = ctx.sqlTx.Query("select * from user_labels where uid=?", uid)
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return
		}
		userLabels[id] = true
	}
	for i, v := range aids {
		rows, err = ctx.sqlTx.Query(`select
			exists(select*from like_answers where uid=? and aid=?)as liked,
			exists(select*from approve_answers where uid=? and aid=?)as approved`,
			uid, v,
			uid, v)
		if err != nil {
			return
		}
		if rows.Next() {
			err = rows.Scan(&infos[i].Liked, &infos[i].Approved)
			if err != nil {
				e := rows.Close()
				if e != nil {
					log.Warn(e)
				}
				return
			}
		}
		e := rows.Close()
		if e != nil {
			log.Warn(e)
		}
		rows, err = ctx.sqlTx.Query("select lid from question_labels where qid=?", qids[i])
		if err != nil {
			return
		}
		infos[i].Approvable = true
		for rows.Next() {
			var id int64
			err = rows.Scan(&id)
			_, ok := userLabels[id]
			if !ok {
				infos[i].Approvable = false
				break
			}
		}
		e = rows.Close()
		if e != nil {
			log.Warn(e)
		}
	}
	return infos, nil
}

func (s *SearchDaoImpl) SearchUsers(ctx TransactionContext, page int64, text string) (result []SearchUserResult, err error) {
	rows, err := ctx.sqlTx.Query("select uid,role,name,nickname,profile from users where match(name,nickname,profile)against(? in boolean mode)limit ?,?",
		text, page*PageSize, PageSize)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var it SearchUserResult
		var role int64
		err = rows.Scan(&it.Uid, &role, &it.Name, &it.Nickname, &it.Profile)
		if err != nil {
			return
		}
		it.Banned = role == entity.DISABLE
		result = append(result, it)
	}
	for i, _ := range result {
		var detail entity.UserDetails
		e := ctx.session.DB("sofia").C("user_details").FindId(result[i].Uid).One(&detail)
		if e != nil {
			result[i].Icon = detail.Icon
		}
	}
	return
}

func (s *SearchDaoImpl) HotList(ctx TransactionContext) (questions []entity.Questions, err error) {
	var rows *sql.Rows
	rows, err = ctx.sqlTx.Query("select "+questionFields+" from questions order by view_count desc limit ?", HotListSize)
	if err != nil {
		return
	}
	defer rows.Close()
	questions, err = s.ParseQuestions(rows)
	if err != nil {
		return
	}
	err = s.AssignLabels(ctx, questions)
	return questions, err
}

func (s *SearchDaoImpl) Search(ctx TransactionContext, text string) (result []KListItem, err error) {
	var rows *sql.Rows
	rows, err = ctx.sqlTx.Query("select kid,title from kcards where match(title)against(?)limit ?", text, KListSize)
	if err != nil {
		return
	}
	result = []KListItem{}
	var ids []int64
	for rows.Next() {
		var it KListItem
		var id int64
		err = rows.Scan(&id, &it.Title)
		if err != nil {
			_ = rows.Close()
			return
		}
		ids = append(ids, id)
		result = append(result, it)
	}
	_ = rows.Close()
	for i, id := range ids {
		rows, err = ctx.sqlTx.Query("select name,value,origin from kcard_attrs where kid=? limit ?", id, KCardSize)
		if err != nil {
			return
		}
		var qid int64
		err = rows.Scan(&result[i].Attr.Name, &result[i].Attr.Value, &qid)
		_ = rows.Close()
		if err != nil {
			return
		}
		result[i].Attr.Origin = strconv.FormatInt(qid, 10)
	}
	return result, nil
}