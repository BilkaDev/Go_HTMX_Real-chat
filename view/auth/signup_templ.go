// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package auth

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/bilkadev/Go_HTMX_Real-chat/view/components/ui"

func SignUp() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col items-center justify-center min-w-96 mx-auto\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 class=\"mb-4 text-3xl font-semibold text-center text-gray-300\">SignUp\r <span class=\"text-blue-500\">GoChat</span></h1><form id=\"login-form\" hx-post=\"/auth/signup\"><div><input type=\"text\" name=\"fullName\" placeholder=\"Enter fullname\" class=\"w-full input input-bordered h-10\" required minlength=\"3\" max=\"100\"></div><div class=\"mt-2\"><input name=\"userName\" value=\"\" type=\"text\" placeholder=\"Enter username\" class=\"w-full input input-bordered h-10\" required minlength=\"3\" maxlength=\"100\"></div><div class=\"mt-2\"><input name=\"email\" type=\"email\" placeholder=\"Enter email\" class=\"w-full input input-bordered h-10\" maxlength=\"255\"></div><div class=\"mt-2\"><input name=\"password\" type=\"password\" placeholder=\"Enter Password\" class=\"w-full input input-bordered h-10\" required minlength=\"6\" max=\"255\"></div><div class=\"mt-2\"><input name=\"repassword\" type=\"password\" placeholder=\"Confirm Password\" class=\"w-full input input-bordered h-10\" required minlength=\"6\" max=\"255\"></div><div class=\"flex\"><div class=\"form-control\"><label class=\"label gap-2 cursor-pointer\"><span class=\"label-text text-indigo-50\">Male</span> <input id=\"gender-male\" type=\"checkbox\" checked name=\"gender\" value=\"male\" class=\"checkbox border-slate-900\"></label></div><div class=\"form-control\"><label class=\"label gap-2 cursor-pointer\"><span class=\"label-text text-indigo-50\">Female</span> <input id=\"gender-female\" type=\"checkbox\" name=\"gender\" value=\"female\" class=\"checkbox border-slate-900\"></label></div></div><a hx-boost=\"true\" href=\"/login\" class=\"text-sm text-indigo-50  hover:underline hover:text-blue-600 mt-2 inline-block\">Already have an account?\r</a><div class=\"mt-2\"><button class=\"btn btn-block btn-sm\"><span className=\"loading loading-spinner \"></span> Sign Up\r</button></div></form>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = ui.GlassCard("w-full p-6 rounded-lg shadow-md").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><script>\r\n\tconst checkboxFemale =\tdocument.getElementById(\"gender-female\")\r\n\tconst checkboxMale =document.getElementById(\"gender-male\")\r\n    \r\n    checkboxFemale.addEventListener('change',(e)=> {\r\n        if (e.target.checked) {\r\n            checkboxMale.checked = false\r\n        }\r\n    })\r\n    \r\n    checkboxMale.addEventListener('change',(e)=> {\r\n        if (e.target.checked) {\r\n            checkboxFemale.checked = false\r\n        }\r\n    })\r\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
