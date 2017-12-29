package text_speech

//
// import (
// 	"errors"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"os"
// 	_ "os/exec"
// 	"os/user"
// 	"strings"
// 	"unicode/utf8"
// )
//
// var langs = []string{"af", "ar", "ca", "cs", "cy", "da", "de", "el", "en",
// 	"es", "fi", "fr", "hi", "hr", "ht", "hu", "hy", "id", "is", "it",
// 	"ja", "ko", "la", "lv", "mk", "nl", "no", "pl", "pt", "ro", "ru",
// 	"sk", "sq", "sr", "sv", "sv", "sw", "ta", "tr", "vi", "zh"}
//
// func Say_t(s string, args ...string) error {
// 	lang := "en"
// 	if len(args) > 0 {
// 		lang = args[0]
// 	}
//
// 	if !isValidLang(lang) {
// 		return errors.New("invalid language code!")
// 	}
//
// 	log.Printf("Lang: %s, say: %s", lang, s)
// 	sentences, err := splitSentences(s)
// 	if err != nil {
// 		log.Print(err)
// 		return err
// 	}
//
// 	err = speak(sentences, lang)
// 	if err != nil {
// 		log.Print(err)
// 		return err
// 	}
//
// 	err = speak(sentences, lang)
// 	if err != nil {
// 		log.Print(err)
// 		return err
// 	}
// 	return nil
//
// }
//
// func isValidLang(s string) bool {
// 	for _, l := range langs {
// 		if l == s {
// 			return true
// 		}
// 	}
// 	log.Printf("Invalid language: %s", s)
// 	return false
// }
//
// func splitSentences(s string) ([]string, error) {
// 	sentences := strings.Split(s, ".")
// 	var result []string
// 	for _, sentence := range sentences {
// 		if utf8.RuneCountInString(sentence) > 100 {
// 			tokens := strings.Split(sentence, ",")
// 			for i, token := range tokens {
// 				tokens[i] = strings.TrimSpace(token)
// 				if utf8.RuneCountInString(tokens[i]) > 100 {
// 					return nil, errors.New("Can't split text into short tokens.")
// 				}
// 			}
// 			result = append(result, tokens...)
// 		} else {
// 			result = append(result, strings.TrimSpace(sentence))
// 		}
// 	}
// 	return result, nil
// }
//
// func getAudio(s, lang string) (io.ReadCloser, error) {
// 	resp, err := http.Get("http://translate.google.com/translate_tts" +
// 		"?ie=UTF-8&tl=" + lang + "&q=" + url.QueryEscape(s))
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return resp.Body, nil
//
// }
//
// func playSound(audio io.Reader, sentences []string, lang string) error {
// 	/*dir, err := os.Open(".")
// 	  if err != nil {
// 	      return err
// 	  }
// 	  defer dir.Close()
// 	  f, err := os.Create("temp.wav")
// 	  if err != nil {
// 	      return err
// 	  }
//
// 	  _, err = io.Copy(f, audio)
// 	  play("temp.wav")*/
// 	// for _, s := range sentences {
// 	// 	filename, err := getFromCache(s, lang)
// 	// 	if err != nil {
// 	// 		return err
// 	// 	}
// 	// 	play(filename)
// 	// }
//
// 	return nil
// }
//
// func getFromCache(s, lang string) (string /*io.ReadCloser*/, error) {
// 	filename := getCacheDir() + "/" + lang + "/" + s + ".mp3"
// 	/*cached, err := os.Open(filename)
// 	return cached, err*/
// 	return filename, nil
// }
//
// func cacheAudio(stream io.Reader, s, lang string) (io.ReadCloser, error) {
// 	langCacheDir := getCacheDir() + "/" + lang
// 	dir, err := os.Open(langCacheDir)
// 	if os.IsNotExist(err) {
// 		_ = os.MkdirAll(langCacheDir, 0700)
// 	}
// 	defer dir.Close()
//
// 	filename := s + "mp3"
//
// 	f, err := os.Open(langCacheDir + "/" + filename)
// 	if os.IsNotExist(err) {
// 		f, err = os.Create(langCacheDir + "/" + filename)
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		_, err = io.Copy(f, stream)
// 		return f, err
// 	}
//
// 	return f, err
// }
//
// func speak(sentences []string, lang string) error {
// 	var streams []io.ReadCloser
// 	for _, s := range sentences {
// 		log.Printf("From cache %s/%s", lang, s)
// 		audio, err := getAudio(s, lang) //getFromCache...
// 		if err != nil {
// 			log.Printf("Cache for %s/%s not found. Trying to get audio from Google", lang, s)
// 			stream, err := getAudio(s, lang)
// 			if err == nil {
// 				log.Printf("Caching stream from %s/%s", lang, s)
// 				audio, _ = cacheAudio(stream, s, lang)
// 			}
// 		}
//
// 		defer audio.Close()
// 		streams = append(streams, audio)
// 	}
//
// 	for _, audio := range streams {
// 		err := playSound(audio, sentences, lang)
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	return nil
// }
//
// func getCacheDir() string {
// 	xdgCacheHome := os.Getenv("XDG_CACHE_HOME")
// 	if xdgCacheHome == "" {
// 		user, _ := user.Current()
// 		home := user.HomeDir
// 		xdgCacheHome = home + "/.chache/gspeak"
// 	}
//
// 	dir, err := os.Open(xdgCacheHome)
// 	if os.IsNotExist(err) {
// 		_ = os.MkdirAll(xdgCacheHome, 0700)
// 		dir, err = os.Open(xdgCacheHome)
// 	}
// 	return dir.Name()
// }
