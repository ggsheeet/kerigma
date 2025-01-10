// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/ggsheet/kerigma/internal/database"
import "github.com/ggsheet/kerigma/template/component"

func Store(books []database.Book, bcategories []database.BCategory, currentPage int, totalPages int) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"es\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\"><title>Tienda Tulip</title><meta name=\"description\" content=\"Tienda Tulip, Editorial Cristiana Reformada\"><meta name=\"robots\" content=\"all\"><link rel=\"icon\" href=\"/public/favicon.ico\" type=\"image/x-icon\"><meta property=\"og:title\" content=\"Tienda Tulip\"><meta property=\"og:description\" content=\"Tienda Tulip, Editorial Cristiana Reformada\"><meta property=\"og:locale\" content=\"es\"><meta property=\"og:site_name\" content=\"Publicaciones Tulip\"><meta property=\"og:url\" content=\"https://www.kerigmalife.com/store\"><meta property=\"og:type\" content=\"website\"><meta property=\"og:image\" content=\"https://kerigmalife.s3.us-east-2.amazonaws.com/og-default.png\"><meta name=\"keywords\" content=\"libreria, libros, cristiana, tienda, reformada, presbiteriana, cristo, dios, leer, blog, articulos, recursos, lectura, teologia\"><link rel=\"canonical\" href=\"https://www.kerigmalife.com/store\"><link rel=\"shortcut icon\" href=\"favicon.ico\"><link rel=\"preload\" href=\"/public/css/reset.css\" as=\"style\"><link rel=\"preload\" href=\"/public/css/global.css\" as=\"style\"><link rel=\"preload\" href=\"/public/css/main.css\" as=\"style\"><link rel=\"preload\" href=\"/public/fonts/DMSans-Regular.woff2\" as=\"font\" type=\"font/woff2\" crossorigin=\"anonymous\"><link rel=\"preload\" href=\"/public/fonts/DMSans-Medium.woff2\" as=\"font\" type=\"font/woff2\" crossorigin=\"anonymous\"><link rel=\"preload\" href=\"/public/fonts/DMSans-Bold.woff2\" as=\"font\" type=\"font/woff2\" crossorigin=\"anonymous\"><link rel=\"stylesheet\" href=\"/public/css/reset.css\"><link rel=\"stylesheet\" href=\"/public/css/global.css\"><link rel=\"stylesheet\" href=\"/public/css/main.css\"><script type=\"text/javascript\" src=\"/public/js/main.js\" defer></script><script type=\"text/javascript\" src=\"/public/js/htmx.min.js\" defer></script></head><body><div id=\"snackBar\" class=\"snack_bar\"><div class=\"snack_text\"></div></div><header class=\"header_styles\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = component.Nav().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</header><main class=\"main_styles\"><h1 class=\"hidden\">Tienda en Línea</h1><div class=\"filters_items\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = component.BookFilter(bcategories).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = component.BookGrid(books, currentPage, totalPages).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></main><footer class=\"footer_styles\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = component.Footer().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</footer></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
