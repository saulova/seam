package middlewares

import (
	"errors"

	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/jwt/configs"

	jwtWare "github.com/gofiber/contrib/jwt"
)

type JWTValidationMiddleware struct{}

const JWTValidationMiddlewareId = "plugins.jwt.middlewares.JWTValidationMiddleware"

func NewJWTValidationMiddleware() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &JWTValidationMiddleware{}

	dependencyContainer.AddDependency(JWTValidationMiddlewareId, instance)

	return instance
}

func (j *JWTValidationMiddleware) createJWTWareConfig(config *configs.JWTMiddlewareConfig) (jwtWare.Config, error) {
	jwtWareConfig := jwtWare.Config{}

	hasSigningKey := len(config.SigningKey.KeyAsByteSlice) > 0
	hasJwkUrls := len(config.JwkUrls) > 0

	if !hasSigningKey && !hasJwkUrls {
		return jwtWareConfig, errors.New("must have signing key or jwk urls")
	}

	if hasSigningKey {
		jwtWareConfig.SigningKey = jwtWare.SigningKey{
			JWTAlg: config.SigningKey.Algorithm,
			Key:    config.SigningKey.KeyAsByteSlice,
		}
	}

	if hasJwkUrls {
		jwtWareConfig.JWKSetURLs = config.JwkUrls
	}

	return jwtWareConfig, nil
}

func (j *JWTValidationMiddleware) Build(config interface{}) (interface{}, error) {
	jwtMiddlewareConfig, err := configs.NewJWTMiddlewareConfig(config)
	if err != nil {
		return nil, err
	}

	jwtConfig, err := j.createJWTWareConfig(jwtMiddlewareConfig)
	if err != nil {
		return nil, err
	}

	middlewareFunc := jwtWare.New(jwtConfig)

	return middlewareFunc, nil
}
