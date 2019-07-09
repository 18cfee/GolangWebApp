
var urlParams = new URLSearchParams(window.location.search);
new Vue({
    el: '#formV1',
    data: {
        edit: false,
        jsonPayload: {
            email: "",
            address: "",
            var3: "",
        },
        formId: urlParams.get('formId'),
        clientId: urlParams.get('userId'),
    },
    created: function() {
        this.initJson();
    },
    methods: {
        editData: function() {
            if (this.edit == true) {
                if (confirm("Do you want to leave edit mode and submit changes?")) {
                    this.edit = false;
                }
            } else {
                this.edit = true;
            }
        },
        resetEdit: function() {
            if (confirm(`Do you want to erase all edits since the edit button was pressed permanently?`)) {
                this.initJson();
            }
        },
        submit: function() {
            axios.put('/updateForm?formId=' + this.formId, this.jsonPayload)
                .then(res => {
                    console.log('it was a success');
                })
                .catch(error => {
                    console.log('something went wrong');
                })
        },
        initJson: function() {
            axios.get('/getFormInfo?formId=' + this.formId)
                .then(res => {
                    console.log("success " + res.data);
                    // 204 is when there was a response but there was no data 
                    if (res.status != 204) {
                        console.log("setting the data");
                        this.jsonPayload = res.data;
                    }
                })
                .catch(error => {
                    console.log("failure");
                })
        },
    }
});
// maybe have the data interactively change on other forms  when someone submits