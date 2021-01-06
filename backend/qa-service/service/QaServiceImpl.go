package service

import (
	"fmt"
	"github.com/SKFE396/qa-service/dao"
	"github.com/SKFE396/qa-service/entity"
	"github.com/SKFE396/qa-service/rpc"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"strconv"
	"strings"
	"time"
)

const (
	ADMIN = 0
	USER  = 1
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
	qaDao    dao.QaDao
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

type ReqAnswersPost struct {
	Qid     string `json:"qid"`
	Content string `json:"content"`
}

type ReqAnswersPut struct {
	Aid     string `json:"aid"`
	Content string `json:"content"`
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
	Time          string   `json:"time"`
	AnswerCount   int64    `json:"answer_count"`
	ViewCount     int64    `json:"view_count"`
	FavoriteCount int64    `json:"favorite_count"`
	Category      string   `json:"category"`
	Labels        []string `json:"labels"`
	Head          string   `json:"head"`
	PictureUrls   []string `json:"pictureUrls"`
}

type QuestionInfo struct {
	Qid           string   `json:"qid"`
	Owner         Owner    `json:"raiser"`
	Title         string   `json:"title"`
	Time          string   `json:"time"`
	AnswerCount   int64    `json:"answer_count"`
	ViewCount     int64    `json:"view_count"`
	FavoriteCount int64    `json:"favorite_count"`
	Category      string   `json:"category"`
	Labels        []string `json:"labels"`
	Content       string   `json:"content"`
	Accepted      string   `json:"accepted_answer"`
}

type AnswerInfo struct {
	Aid            string `json:"aid"`
	Owner          Owner  `json:"answerer"`
	Time           string `json:"time"`
	LikeCount      int64  `json:"like_count"`
	CriticismCount int64  `json:"criticism_count"`
	ApprovalCount  int64  `json:"approval_count"`
	CommentCount   int64  `json:"comment_count"`
	Content        string `json:"content"`
	Liked          bool   `json:"liked"`
	Approved       bool   `json:"approved"`
	Approvable     bool   `json:"approvale"`
}

type AnswerListItem struct {
	Aid            string   `json:"aid"`
	Owner          Owner    `json:"answerer"`
	LikeCount      int64    `json:"like_count"`
	CriticismCount int64    `json:"criticism_count"`
	ApprovalCount  int64    `json:"approval_count"`
	CommentCount   int64    `json:"comment_count"`
	Head           string   `json:"head"`
	Time           string   `json:"time"`
	PictureUrls    []string `json:"picture_urls"`
	Liked          bool     `json:"liked"`
	Approved       bool     `json:"approved"`
	Approvable     bool     `json:"approvale"`
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
		res[i].Time = fmt.Sprint(time.Unix(v.Time, 0))
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
	for i := range res {
		res[i].Owner.Uid = strconv.FormatInt(uids[i], 10)
		res[i].Owner.Name = userInfos[i].Name
		res[i].Owner.Icon = userInfos[i].Icon
		res[i].Owner.Nickname = userInfos[i].Nickname
	}
	return res, nil
}

func (q *QaServiceImpl) AnswerListResponse(ctx dao.TransactionContext, uid int64, answers []entity.Answers, answerDetails []entity.AnswerDetails) (result interface{}, err error) {
	res := make([]AnswerListItem, len(answers))
	uids := make([]int64, len(answers))
	qids := make([]int64, len(answers))
	aids := make([]int64, len(answers))
	for i, v := range answers {
		uids[i] = v.Answerer
		qids[i] = v.Qid
		aids[i] = v.Aid
		res[i].Aid = strconv.FormatInt(v.Aid, 10)
		res[i].LikeCount = v.LikeCount
		res[i].CriticismCount = v.CriticismCount
		res[i].ApprovalCount = v.ApprovalCount
		res[i].CommentCount = v.CommentCount
		res[i].Head = answerDetails[i].Head
		res[i].Time = fmt.Sprint(time.Unix(v.Time, 0))
		if answerDetails[i].PictureUrl != "" {
			res[i].PictureUrls = []string{answerDetails[i].PictureUrl}
		}
	}
	var userInfos []rpc.UserInfo
	userInfos, err = q.usersRPC.GetUserInfos(uids)
	if err != nil {
		return
	}
	for i := range res {
		res[i].Owner.Uid = strconv.FormatInt(uids[i], 10)
		res[i].Owner.Name = userInfos[i].Name
		res[i].Owner.Icon = userInfos[i].Icon
		res[i].Owner.Nickname = userInfos[i].Nickname
	}
	actionInfos, err := q.qaDao.GetAnswerActionInfos(ctx, uid, qids, aids)
	if err != nil {
		return
	}
	for i := range res {
		res[i].Liked = actionInfos[i].Liked
		res[i].Approved = actionInfos[i].Approved
		res[i].Approvable = actionInfos[i].Approvable
	}
	return res, nil
}

func MatchKeywords(text *string, words *[]string) bool {
	str := strings.ToLower(*text)
	for _, v := range *words {
		if strings.Index(str, strings.ToLower(v)) != -1 {
			return true
		}
	}
	return false
}

func FindTextAndPicture(words *[]string, picture *string, headText *[]rune, node *html.Node, findPicture bool, findHead bool) (foundPicture bool, foundHead bool, hasKeywords bool) {
	foundPicture = false
	foundHead = false
	hasKeywords = false
	switch node.Type {
	case html.TextNode:
		data := strings.ReplaceAll(node.Data, "\n", "")
		if findHead && data != "" {
			*headText = append(*headText, []rune(data+" ")...)
			if len(*headText) >= HeadLengthMax {
				foundHead = true
			}
		}
		if MatchKeywords(&node.Data, words) {
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

func (q *QaServiceImpl) ParseContent(content *string, words *[]string) (pictureUrl string, head string, hasKeywords bool) {
	output := blackfriday.Run([]byte(*content))
	output = bluemonday.UGCPolicy().SanitizeBytes(output)
	node, err := html.Parse(strings.NewReader(string(output)))
	if err != nil {
		text := []rune(*content)
		if len(text) > HeadLengthMax {
			return "", string(text[:HeadLengthMax]), MatchKeywords(content, words)
		}
		return "", *content, MatchKeywords(content, words)
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
		return Failed, map[string]int8{"type": UnknownError}
	}
	// serve
	pictureUrl, head, hasKeyword := q.ParseContent(&content, &words)
	if hasKeyword {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
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
	err = q.qaDao.IncQuestionCount(ctx, uid)
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
	pictureUrl, head, hasKeyword := q.ParseContent(&content, &words)
	if hasKeyword {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
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

func (q *QaServiceImpl) AddAnswer(token string, req ReqAnswersPost) (int8, interface{}) {
	content := req.Content
	qid, err := strconv.ParseInt(req.Qid, 10, 64)
	if err != nil {
		return Failed, nil
	}
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
		return Failed, map[string]int8{"type": UnknownError}
	}
	// serve
	pictureUrl, head, hasKeyword := q.ParseContent(&content, &words)
	if hasKeyword {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		return Failed, map[string]int8{"type": HasKeyword}
	}
	aid, err := q.qaDao.AddAnswer(ctx, uid, qid, content, pictureUrl, head)
	if err != nil {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		log.Warn(err)
		return Failed, map[string]int8{"type": UnknownError}
	}
	err = q.qaDao.IncUserAnswerCount(ctx, uid)
	if err != nil {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		log.Warn(err)
		return Failed, map[string]int8{"type": UnknownError}
	}
	question, err := q.qaDao.FindQuestionById(ctx, qid)
	if err != nil {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		log.Warn(err)
		return Failed, map[string]int8{"type": UnknownError}
	}
	if len(question) < 1 {
		log.Warn("AddAnswer: qid = ", qid, ", not found")
	}
	question[0].AnswerCount++
	err = q.qaDao.SaveQuestionSkeleton(ctx, question[0])
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
	return Succeeded, map[string]string{"aid": strconv.FormatInt(aid, 10)}
}

func (q *QaServiceImpl) ModifyAnswer(token string, req ReqAnswersPut) (int8, interface{}) {
	const (
		ConstraintsViolated = 0
		HasKeyword          = 1
		UnknownError        = 2
	)
	aid, err := strconv.ParseInt(req.Aid, 10, 64)
	if err != nil {
		return Failed, map[string]int8{"type": UnknownError}
	}
	content := req.Content
	// check token
	suc, uid, role := q.usersRPC.ParseToken(token)
	if !suc {
		return Expired, nil
	}
	// check constraints
	if len(content) > QuestionContentLengthMax {
		return Failed, map[string]int8{"type": ConstraintsViolated}
	}
	// check authorization
	ctx, err := q.qaDao.Begin(false)
	answer, err := q.qaDao.FindAnswerById(ctx, aid)
	if err != nil {
		_ = q.qaDao.Rollback(&ctx)
		log.Warn(err)
		return Failed, map[string]int8{"type": UnknownError}
	}
	if len(answer) < 1 {
		log.Warn("ModifyAnswer: aid = ", aid, ", not found")
		_ = q.qaDao.Rollback(&ctx)
		log.Warn(err)
		return Failed, map[string]int8{"type": UnknownError}
	}
	if role != ADMIN {
		owner := answer[0].Answerer == uid
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
	pictureUrl, head, hasKeyword := q.ParseContent(&content, &words)
	if hasKeyword {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		return Failed, map[string]int8{"type": HasKeyword}
	}
	err = q.qaDao.ModifyAnswer(ctx, aid, content, pictureUrl, head)
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
	questionDetails := q.qaDao.FindQuestionDetails(ctx, questions)
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

func (q *QaServiceImpl) QuestionDetail(token string, qid int64) (int8, interface{}) {
	// check token
	suc, _, _ := q.usersRPC.ParseToken(token)
	if !suc {
		return Expired, nil
	}
	// serve
	ctx, err := q.qaDao.Begin(true)
	if err != nil {
		return Failed, nil
	}
	question, err := q.qaDao.FindQuestionById(ctx, qid)
	if err != nil {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		log.Warn(err)
		return Failed, nil
	}
	if len(question) < 1 {
		log.Warn("QuestionDetail: qid = ", qid, ", not found")
		return Failed, nil
	}
	detail := q.qaDao.FindQuestionDetails(ctx, question)
	qs := question[0]
	var res QuestionInfo
	res.Qid = strconv.FormatInt(qs.Qid, 10)
	res.Title = qs.Title
	res.Time = fmt.Sprint(qs.Time)
	res.AnswerCount = qs.AnswerCount
	res.ViewCount = qs.ViewCount
	res.FavoriteCount = qs.FavoriteCount
	res.Category = qs.Category
	res.Labels = qs.Labels
	res.Content = detail[0].Content
	if qs.AcceptedAnswer.Valid {
		val, e := qs.AcceptedAnswer.Value()
		if e == nil {
			res.Accepted = strconv.FormatInt(val.(int64), 10)
		}
	}
	uids := []int64{qs.Raiser}
	var userInfos []rpc.UserInfo
	userInfos, err = q.usersRPC.GetUserInfos(uids)
	if err != nil {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		return Failed, nil
	}
	question[0].ViewCount++
	_ = q.qaDao.SaveQuestionSkeleton(ctx, question[0])
	err = q.qaDao.Commit(&ctx)
	if err != nil {
		log.Warn(err)
	}
	res.Owner.Uid = strconv.FormatInt(uids[0], 10)
	res.Owner.Name = userInfos[0].Name
	res.Owner.Icon = userInfos[0].Icon
	res.Owner.Nickname = userInfos[0].Nickname
	return Succeeded, res
}

func (q *QaServiceImpl) ListAnswers(token string, qid int64, page int64, sort int8) (int8, interface{}) {
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
	answers, err := q.qaDao.FindQuestionAnswers(ctx, qid, page, sort)
	if err != nil {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		log.Warn(err)
		return Failed, nil
	}
	answerDetails := q.qaDao.FindAnswerDetails(ctx, answers)
	var result interface{}
	// construct response
	result, err = q.AnswerListResponse(ctx, uid, answers, answerDetails)
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

func (q *QaServiceImpl) AnswerDetail(token string, aid int64) (int8, interface{}) {
	// check token
	suc, uid, _ := q.usersRPC.ParseToken(token)
	if !suc {
		return Expired, nil
	}
	// serve
	ctx, err := q.qaDao.Begin(true)
	if err != nil {
		return Failed, nil
	}
	answer, err := q.qaDao.FindAnswerById(ctx, aid)
	if err != nil {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		log.Warn(err)
		return Failed, nil
	}
	if len(answer) < 1 {
		return Failed, nil
	}
	detail := q.qaDao.FindAnswerDetails(ctx, answer)
	ans := answer[0]
	var res AnswerInfo
	res.Aid = strconv.FormatInt(ans.Aid, 10)
	res.Time = fmt.Sprint(ans.Time)
	res.LikeCount = ans.LikeCount
	res.CriticismCount = ans.CriticismCount
	res.ApprovalCount = ans.ApprovalCount
	res.CommentCount = ans.CommentCount
	res.Content = detail[0].Content

	uids := []int64{ans.Answerer}
	var userInfos []rpc.UserInfo
	userInfos, err = q.usersRPC.GetUserInfos(uids)
	if err != nil {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		log.Warn(err)
		return Failed, nil
	}
	actionInfos, err := q.qaDao.GetAnswerActionInfos(ctx, uid, []int64{ans.Qid}, []int64{aid})
	if err != nil {
		e := q.qaDao.Rollback(&ctx)
		if e != nil {
			log.Warn(e)
		}
		log.Warn(err)
		return Failed, nil
	}
	res.Liked = actionInfos[0].Liked
	res.Approved = actionInfos[0].Approved
	res.Approvable = actionInfos[0].Approvable

	err = q.qaDao.Commit(&ctx)
	if err != nil {
		log.Warn(err)
	}
	res.Owner.Uid = strconv.FormatInt(uids[0], 10)
	res.Owner.Name = userInfos[0].Name
	res.Owner.Icon = userInfos[0].Icon
	res.Owner.Nickname = userInfos[0].Nickname
	return Succeeded, res
}
