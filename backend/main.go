package main

import (
	/* "encoding/xml"    */          // Importa XML apenas por enquanto; o sublinhado impede o uso direto do pacote.
	"fmt"                            // Permite formatar e imprimir mensagens no terminal.
	"github.com/PuerkitoBio/goquery" // Permite pesquisar elementos dentro do HTML.
	"log"                            // Permite registrar erros e encerrar o programa quando necessário.
	"net/http"
	"os" // Permite fazer requisições HTTP para acessar a página do INPI.
	"io"

)

type listXml struct { // Declara uma estrutura para representar dados do XML.
	Processo string `xml:"processo,attr"` // Guarda o atributo processo como texto.
	Titular  string `xml:"titular,attr"`  // Guarda o atributo titular como texto.
	Status   string `xml:"status"`        // Guarda o texto encontrado dentro da tag status.
}

func main() {

	url, err := getScrapeRevista() // Chama a função que acessa e analisa a página da revista.
	if err != nil {
		log.Fatal(err)
	}

	

	err = baixarZip(url)
	if err != nil {
		log.Fatal(err)
	}


	fmt.Printf("URL encontrada: %v\n", url)

} // Finaliza a função principal.

func getScrapeRevista() (string, error) { // Declara uma função para fazer o scraping da página.

	res, err := http.Get("https://revistas.inpi.gov.br/rpi/") // Envia uma requisição GET para a página do INPI.

	if err != nil { // Verifica se ocorreu algum erro ao acessar a URL.
		log.Fatal(err) // Registra o erro e encerra o programa.
	}

	// Finaliza a verificação do erro da requisição.
	defer res.Body.Close() // Agenda o fechamento do conteúdo da resposta quando a função terminar.

	if res.StatusCode != 200 { // Verifica se o servidor respondeu com o código HTTP de sucesso.
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status) // Mostra o código e a descrição do erro.
	} // Finaliza a verificação do status HTTP.

	// Cria um documento pesquisável usando o HTML recebido na resposta.
	doc, err := goquery.NewDocumentFromReader(res.Body) // Converte o corpo da resposta em um documento do goquery.
	if err != nil {                                     // Verifica se houve erro ao interpretar o HTML.
		log.Fatal(err) // Registra o erro e encerra o programa.
	} // Finaliza a verificação do HTML.

	// Procura na página os elementos que correspondem ao seletor CSS informado.

	urlMaisRecente := ""

	doc.Find("tr.warning").Each(func(i int, s *goquery.Selection) {

		link := s.Find("a[href*='/txt/RM']").First()

		url := link.AttrOr("href", "")

		if urlMaisRecente == "" && url != "" {
			urlMaisRecente = url
		}

	}) // Finaliza o percurso dos elementos encontrados.

	return urlMaisRecente, nil // Retorna a URL encontrada e um valor nulo para o erro, indicando sucesso.

} // Finaliza a função de scraping.

func baixarZip(url string) error {

	res, err := http.Get(url) // Envia uma requisição GET para a URL do arquivo ZIP.

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer res.Body.Close() // Agenda o fechamento do conteúdo da resposta quando a função terminar.
	
	arquivo, err := os.Create("rpi.zip") // Cria um arquivo local para salvar o conteúdo do ZIP.

	io.Copy( arquivo, res.Body)

	if err != nil {
		log.Fatal(err)
		return err
	}



	fmt.Println("test", url)

	return nil
}
