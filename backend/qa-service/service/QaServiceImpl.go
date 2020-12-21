package service

import (
	"github.com/SKFE396/qa-service/dao"
	"github.com/SKFE396/qa-service/entity"
	"github.com/SKFE396/qa-service/rpc"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

const (
	ADMIN = iota
	USER
	DISABLE
	NOTACTIVE
)

const (
	Succeeded = iota
	Failed    = iota
	Expired   = iota
)

const (
	HeadLengthMax            = 100
	QuestionLabelsMax        = 5
	QuestionTitleLengthMax   = 32
	LabelLengthMax           = 32
	QuestionContentLengthMax = 50000
)

type QaServiceImpl struct {
	qaDao dao.QaDao
	usersRPC rpc.UsersRPC
}

type ReqQuestionsPost struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Labels   []string `json:"labels"`
}

type ReqQuestionsPut struct {
	Qid      string   `json:"qid"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Labels   []string `json:"labels"`
}

type Owner struct {
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
}

type QuestionListItem struct {
	Qid           string   `json:"qid"`
	Owner         Owner    `json:"raiser"`
	Title         string   `json:"title"`
	Time          int64    `json:"time"`
	AnswerCount   int64    `json:"answer_count"`
	ViewCount     int64    `json:"view_count"`
	FavoriteCount int64    `json:"favorite_count"`
	Category      string   `json:"category"`
	Labels        []string `json:"labels"`
	Head          string   `json:"head"`
	PictureUrls   []string `json:"pictureUrls"`
}

func (q *QaServiceImpl) Init(qaDao dao.QaDao, usersRPC rpc.UsersRPC) (err error) {
	if usersRPC != nil {
		q.usersRPC = usersRPC
	} else {
		q.usersRPC = &rpc.UsersRPCImpl{}
	}
	if qaDao != nil {
		q.qaDao = qaDao
	} else {
		q.qaDao = &dao.QaDaoImpl{}
	}
	return q.qaDao.Init()
}

func (q *QaServiceImpl) Destruct() {
	q.qaDao.Destruct()
}

/*func (q *QaServiceImpl) FindAnswersByQid(qid bson.ObjectId) (answers []entity.Answers, err error) {
	return q.qaDao.FindAnswersByQid(qid)
}

func (q *QaServiceImpl) FindLabelByLid(lid int64) (label entity.Labels, err error) {
	return q.qaDao.FindLabelByLid(lid)
}

func (q *QaServiceImpl) FindLabelByTitle(title string) (label entity.Labels, err error) {
	return q.qaDao.FindLabelByTitle(title)
}

func (q *QaServiceImpl) FindQuestionByQid(qid bson.ObjectId) (question entity.Questions, err error) {
	return q.qaDao.FindQuestionByQid(qid)
}

func (q *QaServiceImpl) FindQuestionLabelsByQid(qid string) (questionLabels []entity.QuestionLabels, err error) {
	return q.qaDao.FindQuestionLabelsByQid(qid)
}

func (q *QaServiceImpl) InsertKcard(kcard entity.Kcards) (kid int64, err error) {
	return q.qaDao.InsertKcard(kcard)
}

func (q *QaServiceImpl) InsertLabel(label entity.Labels) (lid int64, err error) {
	return q.qaDao.InsertLabel(label)
}

func (q *QaServiceImpl) InsertQuestion(question entity.Questions) (qid bson.ObjectId, err error) {
	return q.qaDao.InsertQuestion(question)
}

func (q *QaServiceImpl) InsertQuestionLabel(questionLabel entity.QuestionLabels) (err error) {
	return q.qaDao.InsertQuestionLabel(questionLabel)
}*/

func (q *QaServiceImpl) QuestionListResponse(questions []entity.Questions, questionDetails []entity.QuestionDetails) (result interface{}, err error) {
	res := make([]QuestionListItem, len(questions))
	uids := make([]int64, len(questions))
	for i, v := range questions {
		uids[i] = v.Raiser
		res[i].Qid = strconv.FormatInt(v.Qid, 10)
		res[i].Title = v.Title
		res[i].Time = v.Time
		res[i].AnswerCount = v.AnswerCount
		res[i].ViewCount = v.ViewCount
		res[i].FavoriteCount = v.FavoriteCount
		res[i].Category = v.Category
		res[i].Labels = v.Labels
		res[i].Head = questionDetails[i].Head
		if questionDetails[i].PictureUrl != "" {
			res[i].PictureUrls = []string{questionDetails[i].PictureUrl}
		}
	}
	var userInfos []rpc.UserInfo
	userInfos, err = q.usersRPC.GetUserInfos(uids)
	if err != nil {
		return
	}
	for i, v := range res {
		v.Owner.Uid = strconv.FormatInt(uids[i], 10)
		v.Owner.Name = userInfos[i].Name
		v.Owner.Icon = userInfos[i].Icon
		v.Owner.Nickname = userInfos[i].Nickname
	}
	return res, nil
}

func MatchKeywords(text string, words []string) bool {
	for _, v := range words {
		if strings.Index(text, v) != -1 {
			return true
		}
	}
	return false
}

func FindTextAndPicture(words []string, picture *string, headText *[]rune, node *html.Node, findPicture bool, findHead bool) (foundPicture bool, foundHead bool, hasKeywords bool) {
	foundPicture = false
	foundHead = false
	hasKeywords = false
	switch node.Type {
	case html.TextNode:
		if findHead {
			*headText = append(*headText, []rune(node.Data+" ")...)
			if len(*headText) >= HeadLengthMax {
				foundHead = true
			}
		}
		if MatchKeywords(node.Data, words) {
			return false, false, true
		}
	case html.ElementNode:
		if findPicture && node.Data == "img" {
			for _, v := range node.Attr {
				if v.Key == "src" {
					*picture = v.Val
					foundPicture = true
					break
				}
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		fp, fh, hk := FindTextAndPicture(words, picture, headText, c, !foundPicture, !foundHead)
		foundPicture = foundPicture || fp
		foundHead = foundHead || fh
		if hk {
			return false, false, true
		}
	}
	return
}

func (q *QaServiceImpl) ParseContent(content string, words []string) (pictureUrl string, head string, hasKeywords bool) {
	output := blackfriday.Run([]byte(content))
	output = bluemonday.UGCPolicy().SanitizeBytes(output)
	node, err := html.Parse(strings.NewReader(string(output)))
	if err != nil {
		text := []rune(content)
		if len(text) > HeadLengthMax {
			return "", string(text[:HeadLengthMax]), MatchKeywords(content, words)
		}
		return "", content, MatchKeywords(content, words)
	}
	var headText []rune
	if _, _, hasKeywords = FindTextAndPicture(words, &pictureUrl, &headText, node, true, true); hasKeywords {
		return "", "", true
	}
	if len(headText) > HeadLengthMax {
		headText = headText[:HeadLengthMax]
	}
	return pictureUrl, string(headText), false
}

func (q *QaServiceImpl) AddQuestion(token string, req ReqQuestionsPost) (int8, interface{}) {
	title, labels, content, category := req.Title, req.Labels, req.Content, req.Category
	const (
		ConstraintsViolated = 0
		HasKeyword          = 1
		UnknownError        = 2
	)
	// check token
	suc, uid, _ := q.usersRPC.ParseToken(token)
	if !suc {
		return Expired, nil
	}
	// check constraints
	if len(title) > QuestionTitleLengthMax {
		return Failed, map[string]int8{"type": ConstraintsViolated}
	}
	if len(labels) > QuestionLabelsMax {
		return Failed, map[string]int8{"type": ConstraintsViolated}
	}
	for _, v := range labels {
		if len(v) > LabelLengthMax {
			return Failed, map[string]int8{"type": ConstraintsViolated}
		}
	}
	if len(content) > QuestionContentLengthMax {
		return Failed, map[string]int8{"type": ConstraintsViolated}
	}
	// get banned words
	ctx, err := q.qaDao.Begin(false)
	words, err := q.qaDao.GetBannedWords(ctx)
	if err != nil {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		log.Warn(err)
		return Failed, map[string]int8{"type": ConstraintsViolated}
	}
	// serve
	pictureUrl, head, hasKeyword := q.ParseContent(content, words)
	if hasKeyword {
		return Failed, map[string]int8{"type": HasKeyword}
	}
	qid, err := q.qaDao.AddQuestion(ctx, uid, title, content, category, labels, pictureUrl, head)
	if err != nil {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		log.Warn(err)
		return Failed, map[string]int8{"type": UnknownError}
	}
	err = q.qaDao.Commit(&ctx)
	if err != nil {
		log.Warn(err)
		return Failed, map[string]int8{"type": UnknownError}
	}
	return Succeeded, map[string]string{"qid": strconv.FormatInt(qid, 10)}
}

func (q *QaServiceImpl) ModifyQuestion(token string, req ReqQuestionsPut) (int8, interface{}) {
	const (
		ConstraintsViolated = 0
		HasKeyword          = 1
		UnknownError        = 2
	)
	qid, err := strconv.ParseInt(req.Qid, 10, 64)
	if err != nil {
		return Failed, map[string]int8{"type": UnknownError}
	}
	title, content, category, labels := req.Title, req.Content, req.Category, req.Labels
	// check token
	suc, uid, role := q.usersRPC.ParseToken(token)
	if !suc {
		return Expired, nil
	}
	// check constraints
	if len(title) > QuestionTitleLengthMax {
		return Failed, map[string]int8{"type": ConstraintsViolated}
	}
	if len(content) > QuestionContentLengthMax {
		return Failed, map[string]int8{"type": ConstraintsViolated}
	}
	if len(labels) > QuestionLabelsMax {
		return Failed, map[string]int8{"type": ConstraintsViolated}
	}
	for _, v := range labels {
		if len(v) > LabelLengthMax {
			return Failed, map[string]int8{"type": ConstraintsViolated}
		}
	}
	// check authorization
	ctx, err := q.qaDao.Begin(false)
	if role != ADMIN {
		owner, err := q.qaDao.CheckQuestionOwner(ctx, qid, uid)
		if err != nil {
			e := q.qaDao.Rollback(&ctx)
			if e != nil {
				log.Warn(e)
			}
			log.Warn(err)
			return Failed, map[string]int8{"type": UnknownError}
		}
		if !owner {
			err = q.qaDao.Rollback(&ctx)
			if err != nil {
				log.Warn(err)
			}
			return Failed, map[string]int8{"type": UnknownError}
		}
	}
	// get banned words
	words, err := q.qaDao.GetBannedWords(ctx)
	if err != nil {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		log.Warn(err)
		return Failed, map[string]int8{"type": UnknownError}
	}
	// serve
	pictureUrl, head, hasKeyword := q.ParseContent(content, words)
	if hasKeyword {
		return Failed, map[string]int8{"type": HasKeyword}
	}
	err = q.qaDao.ModifyQuestion(ctx, qid, title, content, category, labels, pictureUrl, head)
	if err != nil {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		log.Warn(err)
		return Failed, map[string]int8{"type": UnknownError}
	}
	err = q.qaDao.Commit(&ctx)
	if err != nil {
		log.Warn(err)
		return Failed, map[string]int8{"type": UnknownError}
	}
	return Succeeded, nil
}

func (q *QaServiceImpl) MainPage(token string, page int64) (int8, interface{}) {
	// check token
	suc, uid, _ := q.usersRPC.ParseToken(token)
	if !suc {
		return Expired, nil
	}
	// check constraints
	if page < 0 {
		return Failed, nil
	}
	// serve
	ctx, err := q.qaDao.Begin(true)
	questions, err := q.qaDao.MainPage(ctx, uid, page)
	if err != nil {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		log.Warn(err)
		return Failed, nil
	}
	questionDetails := q.qaDao.FindDetails(ctx, questions)
	var result interface{}
	// construct response
	result, err = q.QuestionListResponse(questions, questionDetails)
	if err != nil {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		return Failed, nil
	}
	err = q.qaDao.Rollback(&ctx)
	if err != nil {
		log.Warn(err)
		return Failed, nil
	}
	return Succeeded, result
}
