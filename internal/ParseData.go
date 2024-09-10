package internal

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// чтобы было число которое загадал - укажи на 1 больше
func Start(amount int) {
	OpenReadyPage()

	select {}

}

func CopyAlgorithm(page rod.Page) {

}

func OpenReadyPage() {
	url := launcher.New().Headless(false).MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://www.fl.ru/projects/")
	page.MustSetViewport(1200, 800, 1, false)
	page.MustWaitLoad()
	chooseCategories(page)
	ChoosingManyCategories(page)
	ChooseFiltres(page) //он ещё и применяет фильтр
}

func chooseCategories(page *rod.Page) {
	page.MustWaitLoad()
	categoryXPath := `/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[1]/div/div`
	categoryElement, err := page.ElementX(categoryXPath)
	if err != nil {
		fmt.Printf("Ошибка: не удалось найти элемент категории: %v\n", err)

	}
	categoryElement.MustClick()
	time.Sleep(2 * time.Second)
	PathCurrentCategory := `/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[1]/div/ul/li[2]`
	ChooseCurrentCategory, err := page.ElementX(PathCurrentCategory)
	if err != nil {
		fmt.Printf("Ошибка: не удалось выбрать текущую категорию: %v\n", err)
	}
	ChooseCurrentCategory.MustClick()

}

func ChoosingManyCategories(page *rod.Page) {
	CategoriesOpenPath := `/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/div/div[2]`
	elements := []string{
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[2]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[6]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[7]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[8]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[9]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[10]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[13]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[14]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[15]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[16]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[17]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[19]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[20]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[21]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[23]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[25]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[26]",
		"/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[2]/div[3]/div/ul/li[27]",
	}
	for _, elementPath := range elements {
		CategoriesOpen1 := page.MustElementX(CategoriesOpenPath)
		CategoriesOpen1.MustClick()

		// Ждем 0.5 секунды
		time.Sleep(450 * time.Millisecond)

		element := page.MustElementX(elementPath)
		element.MustClick()

		// Ждем 0.5 секунды
		time.Sleep(450 * time.Millisecond)
	}

}

func ChooseFiltres(page *rod.Page) {
	parentElement := page.MustElementX("/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[13]/label")

	parentElement.MustClick()

	time.Sleep(1 * time.Second)

	NoFreelancer := `/html/body/div[2]/div/div[2]/div[2]/div[1]/div/div[14]/label`
	Element2 := page.MustElementX(NoFreelancer)
	Element2.MustClick()
	GoFilterPath := "/html/body/div[2]/div/div[2]/div[2]/div[1]/div/button[1]"
	GoElement := page.MustElementX(GoFilterPath)
	GoElement.MustClick()
	time.Sleep(3 * time.Second) //ожидаем загрузки страницы
}
