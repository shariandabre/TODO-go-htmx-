// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Index() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex-1 h-full w-full flex items-center justify-center bg-teal-lightest font-sans\"><div class=\"bg-white h-full flex flex-col border border-gray-200 rounded-lg shadow-md p-6 m-4 w-full lg:w-3/4 lg:max-w-lg\"><div class=\"mb-4\"><h1 class=\"text-gray-darkest text-center text-2xl\">Todo List</h1><div class=\"flex mt-4\"><form class=\"flex w-full\" hx-post=\"/add-item\" hx-target=\"#item\" hx-swap=\"afterbegin\"><input name=\"add_input\" class=\" shadow appearance-none outline-none border rounded-lg w-full py-2 px-3 mr-4 text-gray-700\" placeholder=\"Add Todo\"> <button type=\"submit\" class=\"flex-no-shrink shadow p-2 border rounded-full text-gray-700 hover:text-blue-500 hover:border-blue-500\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"w-6 h-6\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M12 4.5v15m7.5-7.5h-15\"></path></svg></button></form></div></div><div class=\"overflow-auto flex-1 h-full\"><div class=\"h-[1px] w-full border rounded-full border-gray-200 mt-2 mb-4 bg-gray-200\"></div><div id=\"item\"></div><div><h1 class=\"text-gray-500 \">Completed items</h1><div id=\"completed_items\"></div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
