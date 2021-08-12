import alert from "./../notification/alert.js";

const template = `
<form method="post" action="/login">
    <div class="column row-cols-1 position-absolute top-50 start-50 translate-middle">
        <div class="col-md-auto"  v-if="error">
            <div class="container">
                <alert-component message="Не верно введен логин или пароль." />
            </div>
        </div>

        <div class="col-md-auto">
            <div class="mb-3 row">
                <label for="login" class="col-sm-3 col-form-label">Логин</label>
                <div class="col-sm-9">
                    <input type="login" class="form-control" id="login" name="login">
                </div>
            </div>
        </div>

        <div class="col-md-auto">
            <div class="mb-3 row">
                <label for="password" class="col-sm-3 col-form-label">Пароль</label>
                <div class="col-sm-9">
                    <input type="password" class="form-control" id="password" name="password">
                </div>
            </div>
        </div>

        <div class="col-md-auto">
            <div class="mb-3 form-check">
                <input type="checkbox" checked class="form-check-input" id="remember" name="remember">
                <label class="form-check-label" for="remember">Запомнить вход</label>
            </div>
        </div>

        <button type="submit" class="btn btn-primary">Submit</button>
    </div>
</form>
`;

export function loginComponent(err = false) {
    return {
        data() {
            return {
                error: false
            }
        },

        mounted() {
            this.error = Boolean(err);
        },

        components: {
            "alert-component": alert
        },

        template: template
    }
};
