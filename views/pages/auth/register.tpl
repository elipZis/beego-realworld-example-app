{{ template "../../layout/layout.tpl" . }}

{{ define "content" }}
    <div class="auth-page">
        <div class="container page">
            <div class="row">
                <div class="col-md-6 offset-md-3 col-xs-12">
                    <h1 class="text-xs-center">Sign up</h1>
                    <p class="text-xs-center">
                        <a href="{{urlfor "web.AuthController.GetLogin"}}">Have an account?</a>
                    </p>

                    {{ with .Errors }}
                        <ul class="error-messages">
                            {{range $i, $formError := .}}
                                <li>{{ $i }}: {{ $formError.Message }}</li>
                            {{end}}
                        </ul>
                    {{ end }}

                    <form action="" method="post">
                        <fieldset class="form-group">
                            <input name="username" class="form-control form-control-lg" type="text"
                                   placeholder="Your name"
                                   value="{{ .User.Username }}"/>
                        </fieldset>
                        <fieldset class="form-group">
                            <input name="email" class="form-control form-control-lg" type="text" placeholder="Email"
                                   value="{{ .User.Email }}">
                        </fieldset>
                        <fieldset class="form-group">
                            <input name="password" class="form-control form-control-lg" type="password"
                                   placeholder="Password">
                        </fieldset>
                        <button class="btn btn-lg btn-primary pull-xs-right">
                            Sign up
                        </button>
                    </form>
                </div>

            </div>
        </div>
    </div>
{{ end }}