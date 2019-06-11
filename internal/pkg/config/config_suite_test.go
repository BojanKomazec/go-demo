package config

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BojanKomazec/go-demo/internal/pkg/dbclient"
)

func TestConfig(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("In package config", func() {
	Describe("function ConnParams", func() {
		It("returns the value of connParams field of the receiver", func() {
			connParamsExpected := dbclient.NewConnParams(
				uuid.New().String(),
				rand.Intn(65535),
				uuid.New().String(),
				uuid.New().String(),
				uuid.New().String(),
			)
			dbConfig := DbConfig{connParams: connParamsExpected}
			connParamsActual := dbConfig.ConnParams()
			Expect(connParamsActual.Host()).To(Equal(connParamsExpected.Host()))
			Expect(connParamsActual.Port()).To(Equal(connParamsExpected.Port()))
			Expect(connParamsActual.DbName()).To(Equal(connParamsExpected.DbName()))
			Expect(connParamsActual.User()).To(Equal(connParamsExpected.User()))
			Expect(connParamsActual.Password()).To(Equal(connParamsExpected.Password()))
		})
	})

	Describe("function newImpl", func() {
		Context("when all required environment variables exist and have valid values", func() {
			var kvr = func(key string) (value string, exists bool) {
				if key == "DB_HOST" || key == "DB_NAME" || key == "DB_USER" || key == "DB_PASSWORD" {
					return uuid.New().String(), true
				}

				if key == "DB_PORT" {
					const maxPort = 65535
					return strconv.Itoa(rand.Intn(65535)), true
				}

				return "", false
			}

			Describe("returns a value which", func() {
				config, _ := newImpl(kvr)

				It("should be of type *Config", func() {
					Expect(config).To(BeAssignableToTypeOf(&Config{}))
				})

				It("should contain field DB of type DbConfig", func() {
					dbConfig := config.DB
					Expect(dbConfig).To(BeAssignableToTypeOf(DbConfig{}))
				})
			})
		})

		Context("when DB_HOST environment variable is missing", func() {
			var kvr = func(key string) (value string, exists bool) {
				if key == "DB_HOST" {
					return "", false
				}

				if key == "DB_NAME" || key == "DB_USER" || key == "DB_PASSWORD" {
					return uuid.New().String(), true
				}

				if key == "DB_PORT" {
					const maxPort = 65535
					return strconv.Itoa(rand.Intn(65535)), true
				}

				return "", false
			}
			It("returns an error", func() {
				_, err := newImpl(kvr)
				Expect(err).NotTo(Equal(nil))
			})
		})

		Context("when DB_HOST environment variable is empty string", func() {
			var kvr = func(key string) (value string, exists bool) {
				if key == "DB_HOST" {
					return "", true
				}

				if key == "DB_NAME" || key == "DB_USER" || key == "DB_PASSWORD" {
					return uuid.New().String(), true
				}

				if key == "DB_PORT" {
					const maxPort = 65535
					return strconv.Itoa(rand.Intn(65535)), true
				}

				return "", false
			}
			It("returns an error", func() {
				_, err := newImpl(kvr)
				Expect(err).NotTo(Equal(nil))
			})
		})

		Context("when DB_PORT environment variable is missing", func() {
			var kvr = func(key string) (value string, exists bool) {

				if key == "DB_HOST" || key == "DB_NAME" || key == "DB_USER" || key == "DB_PASSWORD" {
					return uuid.New().String(), true
				}

				if key == "DB_PORT" {
					return "", false
				}

				return "", false
			}

			It("returns an error", func() {
				_, err := newImpl(kvr)
				Expect(err).NotTo(Equal(nil))
			})
		})

		Context("when DB_PORT environment variable is empty string", func() {
			var kvr = func(key string) (value string, exists bool) {

				if key == "DB_HOST" || key == "DB_NAME" || key == "DB_USER" || key == "DB_PASSWORD" {
					return uuid.New().String(), true
				}

				if key == "DB_PORT" {
					return "", true
				}

				return "", false
			}
			It("returns an error", func() {
				_, err := newImpl(kvr)
				Expect(err).NotTo(Equal(nil))
			})
		})

		Context("when DB_NAME environment variable is missing", func() {
			var kvr = func(key string) (value string, exists bool) {
				if key == "DB_NAME" {
					return "", false
				}

				if key == "DB_HOST" || key == "DB_USER" || key == "DB_PASSWORD" {
					return uuid.New().String(), true
				}

				if key == "DB_PORT" {
					const maxPort = 65535
					return strconv.Itoa(rand.Intn(65535)), true
				}

				return "", false
			}
			It("returns an error", func() {
				_, err := newImpl(kvr)
				Expect(err).NotTo(Equal(nil))
			})
		})

		Context("when DB_NAME environment variable is empty string", func() {
			var kvr = func(key string) (value string, exists bool) {
				if key == "DB_NAME" {
					return "", true
				}

				if key == "DB_HOST" || key == "DB_USER" || key == "DB_PASSWORD" {
					return uuid.New().String(), true
				}

				if key == "DB_PORT" {
					const maxPort = 65535
					return strconv.Itoa(rand.Intn(65535)), true
				}

				return "", false
			}
			It("returns an error", func() {
				_, err := newImpl(kvr)
				Expect(err).NotTo(Equal(nil))
			})
		})

		Context("when DB_USER environment variable is missing", func() {
			var kvr = func(key string) (value string, exists bool) {
				if key == "DB_USER" {
					return "", false
				}

				if key == "DB_HOST" || key == "DB_NAME" || key == "DB_PASSWORD" {
					return uuid.New().String(), true
				}

				if key == "DB_PORT" {
					const maxPort = 65535
					return strconv.Itoa(rand.Intn(65535)), true
				}

				return "", false
			}
			It("returns an error", func() {
				_, err := newImpl(kvr)
				Expect(err).NotTo(Equal(nil))
			})
		})

		Context("when DB_USER environment variable is empty string", func() {
			var kvr = func(key string) (value string, exists bool) {
				if key == "DB_USER" {
					return "", true
				}

				if key == "DB_HOST" || key == "DB_NAME" || key == "DB_PASSWORD" {
					return uuid.New().String(), true
				}

				if key == "DB_PORT" {
					const maxPort = 65535
					return strconv.Itoa(rand.Intn(65535)), true
				}

				return "", false
			}
			It("returns an error", func() {
				_, err := newImpl(kvr)
				Expect(err).NotTo(Equal(nil))
			})
		})

		Context("when DB_PASSWORD environment variable is missing", func() {
			var kvr = func(key string) (value string, exists bool) {
				if key == "DB_PASSWORD" {
					return "", false
				}

				if key == "DB_HOST" || key == "DB_NAME" || key == "DB_USER" {
					return uuid.New().String(), true
				}

				if key == "DB_PORT" {
					const maxPort = 65535
					return strconv.Itoa(rand.Intn(65535)), true
				}

				return "", false
			}
			It("returns an error", func() {
				_, err := newImpl(kvr)
				Expect(err).NotTo(Equal(nil))
			})
		})

		Context("when DB_PASSWORD environment variable is empty string", func() {
			var kvr = func(key string) (value string, exists bool) {
				if key == "DB_PASSWORD" {
					return "", true
				}

				if key == "DB_HOST" || key == "DB_NAME" || key == "DB_USER" {
					return uuid.New().String(), true
				}

				if key == "DB_PORT" {
					const maxPort = 65535
					return strconv.Itoa(rand.Intn(65535)), true
				}

				return "", false
			}
			It("returns an error", func() {
				_, err := newImpl(kvr)
				Expect(err).NotTo(Equal(nil))
			})
		})
	})

	Describe("function getEnv", func() {
		Context("when environment variable exists and has some value set", func() {
			expectedValue := uuid.New().String()
			r := func(key string) (value string, exists bool) {
				return expectedValue, true
			}

			It("returns value read from the environment variable", func() {
				actualValue, _ := getEnv(uuid.New().String(), r)
				Expect(actualValue).To(Equal(expectedValue))
			})
		})

		Context("when environment variable does not exist", func() {
			r := func(key string) (value string, exists bool) {
				return "", false
			}

			It("should return a non-nil error", func() {
				_, err := getEnv(uuid.New().String(), r)
				Expect(err).NotTo(Equal(nil))
			})
		})
	})

	Describe("function getEnvAsInt", func() {
		Context("when environment variable exists", func() {
			Context("and has a value that can be converted to an integer", func() {
				expectedValue := rand.Intn(math.MaxUint32)
				expectedValueStr := strconv.Itoa(expectedValue)
				r := func(key string) (value string, exists bool) {
					return expectedValueStr, true
				}

				It("returns an integer created by converting a value read from the environment variable", func() {
					actualValue, _ := getEnvAsInt(uuid.New().String(), r)
					Expect(actualValue).To(Equal(expectedValue))
				})
			})

			Context("and has a value that can't be converted to an integer", func() {
				expectedValueStr := uuid.New().String()
				r := func(key string) (value string, exists bool) {
					return expectedValueStr, true
				}

				It("returns a non-nil error", func() {
					_, err := getEnvAsInt(uuid.New().String(), r)
					Expect(err).NotTo(Equal(nil))
				})
			})
		})

		Context("when environment variable does not exist", func() {
			r := func(key string) (value string, exists bool) {
				return "", false
			}

			It("returns a non-nil error", func() {
				_, err := getEnvAsInt(uuid.New().String(), r)
				Expect(err).NotTo(Equal(nil))
			})
		})
	})

	Describe("function getStringValue", func() {
		When("environment variable doesn't exist", func() {
			kvr := func(key string) (value string, exists bool) {
				return "", false
			}

			It("returns an error", func() {
				_, err := getStringValue(uuid.New().String(), kvr)
				Expect(err).NotTo(Equal(nil))
			})
		})

		When("envirionment variable exists", func() {
			When("its value is empty string", func() {
				kvr := func(key string) (value string, exists bool) {
					return "", true
				}
				It("returns error with a message", func() {
					envVarName := uuid.New().String()
					_, err := getStringValue(envVarName, kvr)
					Expect(err).To(Equal(fmt.Errorf("%s value is empty string", envVarName)))
				})
			})

			When("its value is a non-empty string", func() {
				expectedValue := uuid.New().String()
				kvr := func(key string) (value string, exists bool) {
					return expectedValue, true
				}
				It("returns that value", func() {
					actualValue, _ := getStringValue(uuid.New().String(), kvr)
					Expect(actualValue).To(Equal(expectedValue))
				})
			})
		})
	})

	Describe("function getPortNumber", func() {
		When("environment variable doesn't exist", func() {
			kvr := func(key string) (value string, exists bool) {
				return "", false
			}

			It("returns an error", func() {
				_, err := getPortNumber(uuid.New().String(), kvr)
				Expect(err).NotTo(Equal(nil))
			})
		})

		When("environment variable exists", func() {
			Context("and its value is empty string", func() {
				kvr := func(key string) (value string, exists bool) {
					return "", true
				}
				It("returns error", func() {
					_, err := getPortNumber(uuid.New().String(), kvr)
					Expect(err).NotTo(Equal(nil))
				})
			})

			Context("and its value is a non-empty string", func() {
				Context("which doesn't represent an integer", func() {
					kvr := func(key string) (value string, exists bool) {
						return uuid.New().String(), true
					}
					It("returns an error", func() {
						_, err := getPortNumber(uuid.New().String(), kvr)
						Expect(err).NotTo(Equal(nil))
					})
				})

				Context("which represents an integer", func() {
					Context("which is not in range [0, 65535]", func() {
						expectedValue := rand.Intn(65535) + 65535
						expectedValueStr := strconv.Itoa(expectedValue)
						kvr := func(key string) (value string, exists bool) {
							return expectedValueStr, true
						}
						It("returns error with a message", func() {
							_, err := getPortNumber(uuid.New().String(), kvr)
							Expect(err).To(Equal(fmt.Errorf("Port is out of valid range")))
						})
					})

					Context("which is within range [0, 65535]", func() {
						expectedValue := rand.Intn(65535)
						expectedValueStr := strconv.Itoa(expectedValue)
						kvr := func(key string) (value string, exists bool) {
							return expectedValueStr, true
						}
						It("returns an integer converted from that string", func() {
							actualValue, _ := getPortNumber(uuid.New().String(), kvr)
							Expect(actualValue).To(Equal(expectedValue))
						})
					})
				})
			})
		})
	})
})
