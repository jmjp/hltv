package hltv

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

const HLTV_URL = "https://www.hltv.org"

// FetchConfig contém configurações para o fetcher
type FetchConfig struct {
	MaxRetries    int
	RetryDelay    time.Duration
	EnableLogging bool
}

// DefaultConfig retorna uma configuração padrão
func DefaultConfig() *FetchConfig {
	return &FetchConfig{
		MaxRetries:    3,
		RetryDelay:    time.Second * 2,
		EnableLogging: true,
	}
}

// JA3Pool contém diferentes fingerprints JA3 para rotação
var JA3Pool = []string{
	// Chrome 120+ Windows
	"771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-17513-21,29-23-24,0",
	// Firefox 121+ Windows
	"771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0",
	// Chrome 119 macOS
	"771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,65281-0-23-35-13-5-18-16-30032-11-10,29-23-24,0",
	// Edge 120+ Windows
	"771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27,29-23-24,0",
	// Safari 17+ macOS
	"771,4865-4866-4867-49196-49195-49200-49199-52393-49162-49161-49172-49171-157-156-61-60-53-47-49160-49170-10,65281-0-23-13-5-18-16-11-51-10,29-23-24-25,0",
}

// UserAgentPool contém diferentes user agents para rotação
var UserAgentPool = []string{
	// Chrome Windows
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36",
	// Firefox Windows
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0",
	// Edge Windows
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0",
	// Chrome macOS
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
	// Safari macOS
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Safari/605.1.15",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Safari/605.1.15",
}

// getRandomJA3 retorna um JA3 fingerprint aleatório
func getRandomJA3() string {
	return JA3Pool[rand.Intn(len(JA3Pool))]
}

// getRandomUserAgent retorna um user agent aleatório
func getRandomUserAgent() string {
	return UserAgentPool[rand.Intn(len(UserAgentPool))]
}

// logError registra erros se o logging estiver habilitado
func logError(config *FetchConfig, message string, err error) {
	if config.EnableLogging {
		if err != nil {
			log.Printf("[HLTV Fetcher] %s: %v", message, err)
		} else {
			log.Printf("[HLTV Fetcher] %s", message)
		}
	}
}

// logInfo registra informações se o logging estiver habilitado
func logInfo(config *FetchConfig, message string) {
	if config.EnableLogging {
		log.Printf("[HLTV Fetcher] %s", message)
	}
}

// Fetch faz uma requisição HTTP com fallback mechanisms
func Fetch(url string) (resp *cycletls.Response, err error) {
	return FetchWithConfig(url, DefaultConfig())
}

// FetchWithConfig faz uma requisição HTTP com configuração personalizada
func FetchWithConfig(url string, config *FetchConfig) (resp *cycletls.Response, err error) {
	if config == nil {
		config = DefaultConfig()
	}

	// Inicializa o gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	var lastErr error

	for attempt := 0; attempt <= config.MaxRetries; attempt++ {
		if attempt > 0 {
			logInfo(config, fmt.Sprintf("Tentativa %d/%d para URL: %s", attempt+1, config.MaxRetries+1, url))
			time.Sleep(config.RetryDelay)
		}

		// Seleciona JA3 e User-Agent aleatórios para cada tentativa
		ja3 := getRandomJA3()
		userAgent := getRandomUserAgent()

		logInfo(config, fmt.Sprintf("Usando JA3: %s...", ja3[:50]))
		logInfo(config, fmt.Sprintf("Usando User-Agent: %s", userAgent))

		client := cycletls.Init()
		response, err := client.Do(HLTV_URL+url, cycletls.Options{
			Body:      "",
			Ja3:       ja3,
			UserAgent: userAgent,
			Headers: map[string]string{
				"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
				"Accept-Language":           "en-US,en;q=0.5",
				"Accept-Encoding":           "gzip, deflate, br",
				"DNT":                       "1",
				"Connection":                "keep-alive",
				"Upgrade-Insecure-Requests": "1",
				"Sec-Fetch-Dest":            "document",
				"Sec-Fetch-Mode":            "navigate",
				"Sec-Fetch-Site":            "none",
				"Sec-Fetch-User":            "?1",
			},
		}, "GET")

		if err != nil {
			lastErr = fmt.Errorf("tentativa %d falhou: %w", attempt+1, err)
			logError(config, fmt.Sprintf("Tentativa %d falhou", attempt+1), err)
			continue
		}

		// Verifica diferentes códigos de status e aplica estratégias específicas
		switch response.Status {
		case 200:
			logInfo(config, fmt.Sprintf("Requisição bem-sucedida na tentativa %d", attempt+1))
			return &response, nil
		case 403:
			lastErr = fmt.Errorf("acesso negado (403) - possível detecção de bot")
			logError(config, "Acesso negado (403) - tentando com diferentes fingerprints", nil)
			// Continua para próxima tentativa com diferentes fingerprints
		case 429:
			lastErr = fmt.Errorf("muitas requisições (429) - rate limit atingido")
			logError(config, "Rate limit atingido (429) - aumentando delay", nil)
			// Aumenta o delay para próxima tentativa
			config.RetryDelay *= 2
		case 503, 502, 504:
			lastErr = fmt.Errorf("servidor indisponível (%d)", response.Status)
			logError(config, fmt.Sprintf("Servidor indisponível (%d)", response.Status), nil)
		default:
			lastErr = fmt.Errorf("status inesperado: %d", response.Status)
			logError(config, fmt.Sprintf("Status inesperado: %d", response.Status), nil)
		}
	}

	// Se todas as tentativas falharam
	finalErr := fmt.Errorf("todas as %d tentativas falharam. Último erro: %w", config.MaxRetries+1, lastErr)
	logError(config, "Todas as tentativas falharam", finalErr)
	return nil, finalErr
}

// FetchWithCustomFingerprint permite usar JA3 e User-Agent específicos
func FetchWithCustomFingerprint(url, ja3, userAgent string) (resp *cycletls.Response, err error) {
	config := DefaultConfig()

	logInfo(config, fmt.Sprintf("Usando fingerprint customizado para URL: %s", url))
	logInfo(config, fmt.Sprintf("JA3 customizado: %s...", ja3[:50]))
	logInfo(config, fmt.Sprintf("User-Agent customizado: %s", userAgent))

	client := cycletls.Init()
	response, err := client.Do(HLTV_URL+url, cycletls.Options{
		Body:      "",
		Ja3:       ja3,
		UserAgent: userAgent,
		Headers: map[string]string{
			"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
			"Accept-Language":           "en-US,en;q=0.5",
			"Accept-Encoding":           "gzip, deflate, br",
			"DNT":                       "1",
			"Connection":                "keep-alive",
			"Upgrade-Insecure-Requests": "1",
		},
	}, "GET")

	if err != nil {
		logError(config, "Requisição com fingerprint customizado falhou", err)
		return nil, err
	}

	if response.Status != 200 {
		err := fmt.Errorf("status inesperado: %d", response.Status)
		logError(config, "Status inesperado com fingerprint customizado", err)
		return nil, err
	}

	logInfo(config, "Requisição com fingerprint customizado bem-sucedida")
	return &response, nil
}
