import navbar from "./components/NavbarComponent.js";

const template = `
<div>
    <navbar :auth="false"></navbar>
    <div class="container">
        <br>
        <p>
            <router-link to="/">Go to Home</router-link>
            <br>
            <router-link class="nav-link active" aria-current="page" to="/about">Go to About</router-link>
        </p>
        <br><hr><br>
        <router-view></router-view>
    </div>
</div>
`

const Home = { template: '<div>Home</div>' };
const About = { template: '<div>About</div>' };

const routes = [
    { path: '/', component: Home },
    { path: '/about', component: About },
];

const router = VueRouter.createRouter({
    history: VueRouter.createWebHashHistory(),
    routes,
});

Vue.createApp({
    components: {
        "navbar": navbar
    },
    template: template
})
    .use(router)
    .mount("#app")
;
