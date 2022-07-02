package initiator

import (
	"github.com/nenecchuu/lizbeth-be-core/config"
	"github.com/nenecchuu/lizbeth-be-core/init/service"
	ah "github.com/nenecchuu/lizbeth-be-core/internal/app/auth/handler"
	chh "github.com/nenecchuu/lizbeth-be-core/internal/app/chore/handler"
)

func (i *Initiator) InitRestHandler(cfg *config.MainConfig, infra *service.Infrastructure, uc *service.Usecases) *service.RestHandlers {
	auth := ah.NewRest(ah.RestOpts{AuthUsecase: uc.AuthUsecase})

	return &service.RestHandlers{
		AuthRestHandler: auth,
	}
}

func (i *Initiator) InitChatbotHandler(cfg *config.MainConfig, infra *service.Infrastructure, integration *service.Integration, uc *service.Usecases) *service.ChatbotHandlers {
	auth := ah.NewChatbot(ah.ChatbotOpts{AuthUsecase: uc.AuthUsecase})
	chore := chh.NewChatbot(chh.ChatbotOpts{ChoreUsecase: uc.ChoreUsecase})

	return &service.ChatbotHandlers{
		AuthChatbotHandler:  auth,
		ChoreChatbotHandler: chore,
	}
}
