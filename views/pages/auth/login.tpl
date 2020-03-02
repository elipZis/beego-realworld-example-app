{{ template "../../layout/layout.tpl" . }}

{{ define "content" }}
    <div class="auth-page">
        <div class="container page">
            <div class="row">
                <div class="col-md-6 offset-md-3 col-xs-12">
                    <h1 class="text-xs-center">Sign in</h1>
                    <p class="text-xs-center">
                        <a href="{{urlfor "web.AuthController.GetRegister"}}">Need an account?</a>
                    </p>

                    <ul class="error-messages">
                        <li>That email is already taken</li>
                    </ul>

                    <form action="" method="post">
                        <fieldset class="form-group">
                            <input name="email" class="form-control form-control-lg" type="text" placeholder="Email">
                        </fieldset>
                        <fieldset class="form-group">
                            <input name="password" class="form-control form-control-lg" type="password"
                                   placeholder="Password">
                        </fieldset>
                        <button class="btn btn-lg btn-primary pull-xs-right">
                            Sign in
                        </button>
                    </form>
                </div>

            </div>
        </div>
    </div>
{{ end }}