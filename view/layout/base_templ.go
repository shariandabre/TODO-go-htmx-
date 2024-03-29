// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Base() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html lang=\"en\"><head><title></title><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><script src=\"https://cdn.tailwindcss.com\"></script><style>\n\t\t\t :root {\n        --code-color: darkred;\n        --code-bg-color: #aaaaaa;\n        --code-font-size: 14px;\n        --code-line-height: 1.4;\n        --scroll-bar-color: #c5c5c5;\n        --scroll-bar-bg-color: #f6f6f6;\n    }\n\n    pre {\n        color: var(--code-color);\n        font-size: var(--code-font-size);\n        line-height: var(--code-line-height);\n        background-color: var(--code-bg-color);\n    }\n\n    .code-block {\n        max-height: 100px;\n        overflow: auto;\n        padding: 8px 7px 5px 15px;\n        margin: 0px 0px 0px 0px;\n        border-radius: 7px;\n    }\n\n    ::-webkit-scrollbar-corner { background: rgba(0,0,0,0.5); }\n\n    * {\n        scrollbar-width: thin;\n        scrollbar-color: var(--scroll-bar-color) var(--scroll-bar-bg-color);\n    }\n\n    /* Works on Chrome, Edge, and Safari */\n    *::-webkit-scrollbar {\n        width: 12px;\n        height: 12px;\n    }\n\n    *::-webkit-scrollbar-track {\n        background: var(--scroll-bar-bg-color);\n    }\n\n    *::-webkit-scrollbar-thumb {\n        background-color: var(--scroll-bar-color);\n        border-radius: 20px;\n        border: 3px solid var(--scroll-bar-bg-color);\n    }\n\t\t\t</style><script src=\"https://unpkg.com/htmx.org@1.9.10\" integrity=\"sha384-D1Kt99CQMDuVetoL1lrYwg5t+9QdHe7NLX/SoJYkXDFfX37iInKRy5xLSi8nO7UC\" crossorigin=\"anonymous\"></script><script src=\"https://unpkg.com/htmx.org/dist/ext/multi-swap.js\"></script></head><body><div class=\"flex-1 h-full p-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Index().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
