package article

type articleService struct {
	articleRepository articleRepositoryInterface
}

func NewArticleService(r articleRepositoryInterface) articleServiceInterface {
	return &articleService{
		articleRepository: r,
	}
}

func (s *articleService) Create(a Article) error {
	a.Slug = a.Title
	return s.articleRepository.CreateArticle(a)
}
