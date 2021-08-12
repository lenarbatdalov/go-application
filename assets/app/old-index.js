import todoItem from "./todo-item.js";

let App = {
    data() {
        return {
            counter: 0,
            message: 'Привет, Vue.js!',
            seen: true,
            todos: [
                { text: 'Learn JavaScript' },
                { text: 'Learn Vue' },
                { text: 'Build something awesome' }
            ]
        }
    },

    mounted() {
        setInterval(() => {
            this.counter++;
        }, 1000)
    },

    components: {
        "todo-item": todoItem
    },

    methods: {
        reverseMessage() {
            this.seen = !this.seen;
            this.message = this.message.split('').reverse().join('');
        }
    },

    template: `
        <div>
            <div>Счётчик: {{ counter }}</div>

            <hr>

            <div>
                <p>{{ message }}</p>
                <button v-on:click="reverseMessage">Перевернуть сообщение</button>
            </div>

            <hr>

            <div>
                <span v-if="seen">Сейчас меня видно</span>
            </div>

            <hr>

            <div>
                <ol>
                    <li v-for="todo in todos">
                    {{ todo.text }}
                    </li>
                </ol>
            </div>

            <hr>

            <todo-item
                v-for="(todo, item) in todos"
                :todo="todo.text"
                :key="item"
            ></todo-item>
        </div>
    `
};

Vue.createApp(App).mount("#app");
