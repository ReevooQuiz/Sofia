package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"search-service/dao"
	"search-service/entity"
	"search-service/rpc"
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

type SearchServiceImpl struct {
	searchDao dao.SearchDao
	usersRPC  rpc.UsersRPC
}

type Owner struct {
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
}

type QuestionListItem struct {
	HasKeywords   bool     `json:"has_keywords"`
	Qid           string   `json:"qid"`
	Closed        bool     `json:"closed"`
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

type AnswerListItem struct {
	HasKeywords    bool     `json:"has_keywords"`
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

func (s *SearchServiceImpl) Init(searchDao dao.SearchDao, usersRPC rpc.UsersRPC) (err error) {
	if usersRPC != nil {
		s.usersRPC = usersRPC
	} else {
		s.usersRPC = &rpc.UsersRPCImpl{}
	}
	if searchDao != nil {
		s.searchDao = searchDao
	} else {
		s.searchDao = &dao.SearchDaoImpl{}
	}
	return s.searchDao.Init()
}

func (s *SearchServiceImpl) Destruct() {
	s.searchDao.Destruct()
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

func (s *SearchServiceImpl) QuestionListResponse(questions []entity.Questions, questionDetails []entity.QuestionDetails, keywords *[]string) (result interface{}, err error) {
	res := make([]QuestionListItem, len(questions))
	uids := make([]int64, len(questions))
	for i, v := range questions {
		uids[i] = v.Raiser
		res[i].Qid = strconv.FormatInt(v.Qid, 10)
		res[i].Closed = v.Closed
		if MatchKeywords(&questionDetails[i].Title, keywords) {
			res[i].Title = "[标题包含敏感词，已屏蔽]"
		} else {
			res[i].Title = questionDetails[i].Title
		}
		res[i].Time = fmt.Sprint(time.Unix(v.Time, 0))
		res[i].AnswerCount = v.AnswerCount
		res[i].ViewCount = v.ViewCount
		res[i].FavoriteCount = v.FavoriteCount
		res[i].Category = v.Category
		res[i].Labels = v.Labels
		res[i].HasKeywords = MatchKeywords(&questionDetails[i].Content, keywords)
		if !res[i].HasKeywords {
			res[i].Head = questionDetails[i].Head
		}
		if questionDetails[i].PictureUrl != "" {
			res[i].PictureUrls = []string{questionDetails[i].PictureUrl}
		}
	}
	var userInfos []rpc.UserInfo
	userInfos, err = s.usersRPC.GetUserInfos(uids)
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

func (s *SearchServiceImpl) SearchQuestions(token string, page int64, text string) (code int8, result interface{}) {
	// check token
	suc, _, _ := s.usersRPC.ParseToken(token)
	if !suc {
		return Expired, nil
	}
	// serve
	ctx, err := s.searchDao.Begin(true)
	if err != nil {
		return Failed, nil
	}
	questions, err := s.searchDao.SearchQuestions(ctx, page, text)
	if err != nil {
		s.searchDao.Rollback(&ctx)
		log.Warn(err)
		return Failed, nil
	}
	details := s.searchDao.FindQuestionDetails(ctx, questions)
	keywords, err := s.searchDao.GetBannedWords(ctx)
	if err != nil {
		s.searchDao.Rollback(&ctx)
		log.Warn(err)
		return Failed, nil
	}
	// construct response
	result, err = s.QuestionListResponse(questions, details, &keywords)
	if err != nil {
		s.searchDao.Rollback(&ctx)
		return Failed, nil
	}
	s.searchDao.Rollback(&ctx)
	return Succeeded, result
}

func (s *SearchServiceImpl) AnswerListResponse(ctx dao.TransactionContext, uid int64, answers []entity.Answers, answerDetails []entity.AnswerDetails, keywords *[]string) (result interface{}, err error) {
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
		res[i].Time = fmt.Sprint(time.Unix(v.Time, 0))
		res[i].HasKeywords = MatchKeywords(&answerDetails[i].Content, keywords)
		if !res[i].HasKeywords {
			res[i].Head = answerDetails[i].Head
		}
		if answerDetails[i].PictureUrl != "" {
			res[i].PictureUrls = []string{answerDetails[i].PictureUrl}
		}
	}
	var userInfos []rpc.UserInfo
	userInfos, err = s.usersRPC.GetUserInfos(uids)
	if err != nil {
		return
	}
	for i := range res {
		res[i].Owner.Uid = strconv.FormatInt(uids[i], 10)
		res[i].Owner.Name = userInfos[i].Name
		res[i].Owner.Icon = userInfos[i].Icon
		res[i].Owner.Nickname = userInfos[i].Nickname
	}
	actionInfos, err := s.searchDao.GetAnswerActionInfos(ctx, uid, qids, aids)
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

func (s *SearchServiceImpl) SearchAnswers(token string, page int64, text string) (code int8, result interface{}) {
	// check token
	suc, uid, _ := s.usersRPC.ParseToken(token)
	if !suc {
		return Expired, nil
	}
	// serve
	ctx, err := s.searchDao.Begin(true)
	if err != nil {
		return Failed, nil
	}
	details, err := s.searchDao.SearchAnswers(ctx, page, text)
	if err != nil {
		s.searchDao.Rollback(&ctx)
		log.Warn(err)
		return Failed, nil
	}
	answers := s.searchDao.FindAnswerSkeletons(ctx, details)
	keywords, err := s.searchDao.GetBannedWords(ctx)
	if err != nil {
		s.searchDao.Rollback(&ctx)
		log.Warn(err)
		return Failed, nil
	}
	// construct response
	result, err = s.AnswerListResponse(ctx, uid, answers, details, &keywords)
	if err != nil {
		s.searchDao.Rollback(&ctx)
		return Failed, nil
	}
	s.searchDao.Rollback(&ctx)
	return Succeeded, result
}

func (s *SearchServiceImpl) SearchUsers(token string, page int64, text string) (code int8, result interface{}) {
	// check token
	suc, _, _ := s.usersRPC.ParseToken(token)
	if !suc {
		return Expired, nil
	}
	// serve
	ctx, err := s.searchDao.Begin(true)
	if err != nil {
		log.Warn(err)
		return Failed, nil
	}
	result, err = s.searchDao.SearchUsers(ctx, page, text)
	if err != nil {
		s.searchDao.Rollback(&ctx)
		log.Warn(err)
		return Failed, nil
	}
	s.searchDao.Rollback(&ctx)
	return Succeeded, result
}

func (s *SearchServiceImpl) HotList(token string) (code int8, result interface{}) {
	// check token
	suc, _, _ := s.usersRPC.ParseToken(token)
	if !suc {
		return Expired, nil
	}
	// serve
	ctx, err := s.searchDao.Begin(true)
	if err != nil {
		log.Warn(err)
		return Failed, nil
	}
	questions, err := s.searchDao.HotList(ctx)
	if err != nil {
		s.searchDao.Rollback(&ctx)
		log.Warn(err)
		return Failed, nil
	}
	details := s.searchDao.FindQuestionDetails(ctx, questions)
	keywords, err := s.searchDao.GetBannedWords(ctx)
	if err != nil {
		s.searchDao.Rollback(&ctx)
		log.Warn(err)
		return Failed, nil
	}
	// construct response
	result, err = s.QuestionListResponse(questions, details, &keywords)
	if err != nil {
		s.searchDao.Rollback(&ctx)
		return Failed, nil
	}
	s.searchDao.Rollback(&ctx)
	return Succeeded, result
}

func (s *SearchServiceImpl) Search(token string, text string) (code int8, result interface{}) {
	// check token
	suc, _, _ := s.usersRPC.ParseToken(token)
	if !suc {
		return Expired, nil
	}
	// serve
	ctx, err := s.searchDao.Begin(true)
	if err != nil {
		log.Warn(err)
		return Failed, nil
	}
	result, err = s.searchDao.Search(ctx, text)
	if err != nil {
		s.searchDao.Rollback(&ctx)
		log.Warn(err)
		return Failed, nil
	}
	s.searchDao.Rollback(&ctx)
	return Succeeded, result
}
