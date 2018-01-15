package dash

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"html/template"
	"log"
	"strings"

	"github.com/gobuffalo/plush"
	"github.com/julienschmidt/httprouter"
	"gobot.io/x/gobot/drivers/gpio"
	"strconv"
)

var rgbled *gpio.RgbLedDriver

var tmplCache map[string]string

func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := plush.NewContext()
	ctx.Set("title", "Robot dashboard ultimate")

	ctx.Set("partial", partial(ctx))
	ctx.Set("url", url(ctx))

	s, err := plush.Render(loadTemplate("main.html"), ctx)

	if err != nil {
		fmt.Println("ERROR", err)
	}
	fmt.Fprint(w, s)

}

func handlerSetRGB(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	red, _ := strconv.Atoi(ps.ByName("r"))
	green, _ := strconv.Atoi(ps.ByName("g"))
	blue, _ := strconv.Atoi(ps.ByName("b"))

	fmt.Println("RGB", red, green, blue)
	rgbled.SetRGB(byte(red), byte(green), byte(blue))
	fmt.Fprint(w, "ok")
}

func partial(ctx *plush.Context) func(string) (template.HTML, error) {
	return func(name string) (template.HTML, error) {
		t, err := plush.Render(loadTemplate(name), ctx)
		return template.HTML(t), err
	}
}

func url(ctx *plush.Context) func(string) (template.HTML, error) {
	return func(str string) (template.HTML, error) {
		str = strings.Trim(str, " ")
		return template.HTML(str), nil
	}
}

func loadTemplate(name string) string {
	content, err := ioutil.ReadFile("./example9/dash/views/" + name)
	fmt.Println("loading:", "./example9/dash/views/"+name)
	if err != nil {
		fmt.Println(err, name)
		return ""
	}
	tmplCache[name] = string(content)
	return string(content)
}

func Start(led *gpio.RgbLedDriver) {
	rgbled = led
	tmplCache = make(map[string]string)

	router := httprouter.New()
	router.GET("/", handler)
	router.GET("/setrgb/:r/:g/:b", handlerSetRGB)
	router.ServeFiles("/static/*filepath", http.Dir("example9/dash/static"))

	go func() {
		fmt.Println("Dashboard up on port 3010")
		log.Fatal(http.ListenAndServe(":3010", router))
		fmt.Println("Dashboard down")
	}()

}
