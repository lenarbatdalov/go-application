export default {
    props: ['todo'],

    mounted() {
        console.log(this.todo);
    },
    
    template: `<li>{{ todo }}</li>`
}
