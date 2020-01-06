<template>
  <div class="hello">
    <!-- <h1>{{ msg }}</h1> -->

    <div class="main">
      <h1 class="title">Deployments</h1>
    </div>
    <hr>

    <div>
      <b-container fluid>
        <div class="deployments">
        <b-row class="align-self-auto">
          <b-col cols="5" class="text-center"></b-col>
          <b-col cols="2" class="text-center">
            <form>

              <div class="field">
                <label class="label">ServiceName: </label>
                <b-form-input name="servicename" v-model="servicename" v-validate="'required|string'" class="input" type="text"></b-form-input>
              </div>
              <div class="field">
                <label class="label">ReplicaCount: </label>
                <b-form-input name="replicacount" v-model="replicacount" v-validate="'required|string'" class="input" type="text"></b-form-input>
              </div>
              <div class="field">
                <label class="label">ImageName: </label>
                <b-form-input name="imagename" v-model="imagename" v-validate="'required|string'" class="input" type="text"></b-form-input>
              </div>
              <div class="field">
                <label class="label">Repo: </label>
                <b-form-input name="repo" v-model="repo" v-validate="'required|string'" class="input" type="text"></b-form-input>
              </div>
              <div class="field">
                <label class="label">ImageTag: </label>
                <b-form-input name="imagetag" v-model="imagetag" v-validate="'required|string'" class="input" type="text"></b-form-input>
              </div>
              <div class="field">
                <label class="label">PullPolicy: </label>
                <b-form-input name="pull" v-model="pull" v-validate="'required|string'" class="input" type="text"></b-form-input>
              </div>
              <div class="field">
                <label class="label">LoadBalancer: </label>
                <b-form-input name="lb" v-model="lb" v-validate="'required|string'" class="input" type="text"></b-form-input>
              </div>
              <div class="field">
                <label class="label">ExternalPort: </label>
                <b-form-input name="externalport" v-model="externalport" v-validate="'required|string'" class="input" type="text"></b-form-input>
              </div>
              <div class="field">
                <label class="label">InternalPort: </label>
                <b-form-input name="internalport" v-model="internalport" v-validate="'required|string'" class="input" type="text"></b-form-input>
              </div>
              <br>
            </form>
          </b-col>
        </b-row>
        </div>
      </b-container>
    </div>
    <br>
    <div><p>these are the deploysments returned {{ deployments }}</p></div>
  <br>
    &nbsp;
    <div class="buttons">
      <b-row class="align-self-auto">
        <b-col cols="5" class="text-center"></b-col>
        <b-col cols="2" class="text-center">
          <b-button variant="primary" v-on:click="postreq()">Deploy</b-button>
          <div class="divider"></div>
          <b-button variant="primary" v-on:click="getDeploy()">Get Deployments</b-button>
        </b-col>
      </b-row>
    </div>

    &nbsp;
    <hr>

  </div>
</template>

<script>

  import axios from 'axios';
  import Vue from 'vue'
  import * as VeeValidate from 'vee-validate';

  /* eslint-disable */
  Vue.use(VeeValidate)

  export default {
    name: 'Deploy',

    data: function () {
      return {
        servicename: "",
        replicacount: "",
        imagename: "",
        repo: "",
        imagetag: "",
        pull: "",
        lb: "",
        externalport: "",
        internalport: ""

      }
    },

    methods: {
      postreq: function () {
        var data = {"servicename": String(this.servicename)}

        /*eslint-disable*/
        console.log(data)
        /*eslint-enable*/

        axios({
          method: "POST",
          url: "http://localhost:3030/api/set",
          data: data,
          headers: {"content-type": "text/plain"}
        }).then(result => {
          // this.response = result.data;

          /*eslint-disable*/
          console.log(result.data)
          /*eslint-enable*/

        }).catch(error => {
          /*eslint-disable*/
          console.error(error);
          /*eslint-enable*/
        });
        window.location.reload()
      },

      data: function() {
        return {
          deployments: []
        }
      },

      getDeploy: function () {
        var deployments = []
        axios({
          method: "GET",
          url: "http://localhost:3030/api/deployments",
        }).then(result => {
          // this.response = result.data;
          deployments = result.data
          /*eslint-disable*/
          console.log(result.data)
          /*eslint-enable*/

        }).catch(error => {
          /*eslint-disable*/
          console.error(error);
          /*eslint-enable*/
        });

        document.getElementById('deployments').innerHTML = deployments;
      }
    }
  }

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  h3 {
    margin: 40px 0 0;
  }
  ul {
    list-style-type: none;
    padding: 0;
  }
  li {
    display: inline-block;
    margin: 0 10px;
  }
  a {
    color: #42b983;
  }

  .main {
    padding-bottom: auto;
    padding-left: auto;
    background-color: whitesmoke;
  }

  .deployments {
    background-color: lightskyblue;
    text-decoration-color: white;
  }

  .buttons {
    position: relative;
    float: inherit;
  }

  .divider{
    width:5px;
    height:auto;
    display:inline-block;
  }
</style>