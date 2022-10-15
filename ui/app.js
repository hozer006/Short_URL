


new Vue({
  el: "#app",
  data() {
    return {
      qrcode: new QRious({ size: 200 }),
      text: "",
      form: {
        redirect: "",
        random: true,
      },
      info: null,
      msg: null,
    };
  },

  methods: {
    mounted() {
      axios
        .get("http://localhost:3000/short")
        .then((response) => (this.info = response.data));
    },

    submit() {
      axios.post("http://localhost:3000/short", this.form).then(
        function (response) {
          // Handle success
        }.bind(this)
      );
    },

      async deleteData(id) {
        let x = window.confirm("You want to delete the user?");
        if (x) {
          const user = await axios.delete('http://localhost:3000/short/'+id);
        }
      },

      test(redirect) {
        this.text = redirect
        console.log(this.text)
      }
      
    },

    computed: {
      newQRCode() {
        this.qrcode.value = this.text;
        return this.qrcode.toDataURL();
      },
    },
  
  

    
});
