// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func AddProductForm() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col justify-center items-center m-auto gap-4\"><form id=\"imageUploadForm\" enctype=\"multipart/form-data\" class=\"flex flex-col justify-center items-center gap-2 rounded-md border border-black\"><label for=\"image\">Image:</label> <input type=\"file\" id=\"image\" name=\"image\" accept=\"image/*\" required> <button type=\"submit\" class=\"bg-gray-300\">Upload Image</button></form><form id=\"productForm\" action=\"/products\" method=\"POST\" class=\"flex flex-col justify-center items-center gap-2 rounded-md border border-black\"><label for=\"name\">Name:</label> <input type=\"text\" id=\"name\" name=\"name\" required> <label for=\"description\">Description:</label> <textarea id=\"description\" name=\"description\" rows=\"4\" cols=\"50\" required></textarea> <label for=\"quantity\">Quantity:</label> <input type=\"number\" id=\"quantity\" name=\"quantity\" required> <label for=\"price\">Price:</label> <input type=\"number\" id=\"price\" name=\"price\" step=\"0.01\" required> <label for=\"category\">Category:</label> <input type=\"text\" id=\"category\" name=\"category\" required> <input type=\"hidden\" id=\"image_url\" name=\"image_url\"> <input type=\"submit\" value=\"Submit Product\" class=\"bg-gray-300\"></form></div><script type=\"text/javascript\">\n        document.getElementById('imageUploadForm').addEventListener('submit', function(e) {\n            e.preventDefault();\n            \n            var formData = new FormData(this);\n            \n            fetch('/upload-image', {\n                method: 'POST',\n                body: formData\n            })\n            .then(response => response.json())\n            .then(data => {\n                document.getElementById('image_url').value = data.url;\n                alert('Image uploaded successfully');\n            })\n            .catch(error => {\n                console.error('Error:', error);\n                alert('Error uploading image');\n            });\n        });\n\n        document.getElementById('productForm').addEventListener('submit', function(e) {\n            if (!document.getElementById('image_url').value) {\n                e.preventDefault();\n                alert('Please upload an image first');\n            }\n        });\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
