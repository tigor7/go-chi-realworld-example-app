package article

type articleService struct {
	articleRepository articleRepositoryInterface
}

func NewArticleService(r articleRepositoryInterface) articleServiceInterface {
	return &articleService{
		articleRepository: r,
	}
}
