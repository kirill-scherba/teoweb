<template>
  <div class="monitor">
    <h1>This is a monitor page</h1>
    <div>{{ dcAnswer }}</div>
  </div>
</template>

<script> 

var msgID = 0;

/* eslint no-undef: "off"*/
export default {
  name: 'Monitor',
  data: function() {
    // return data
    return {
      dcAnswer:  "Wait data...",
      num:  0,
    }
  },  
  created() {
    // this.interval = setInterval(this.refreshData, 500)
    this.interval2 = setInterval(this.sendData, 5000)
    this.initCallback()
  },
  beforeDestroy() {
    clearInterval(this.interval)
    clearInterval(this.interval2)
  },
  methods: {
    // refreshData () {
    //   this.dcAnswer = data.dcAnswer
    // },
    sendData () {
      var cmd = 129
      SendCmd(cmd, "Hello " + ++msgID + "!", (err, data) => {
        if (err) {
          console.error(err);
          return;
        }
        console.log("Got answer to SendCmd, cmd:", cmd, "data:", data)
        this.dcAnswer = data
      })
      // async function doStuff() {
      //   let data = await SendCmdPromise(cmd, "Hello " + ++msgID + "!")
      //   console.log("Got answer to SendCmd, cmd:", cmd, "data:", data)
      //   this.dcAnswer = data
      // }
      // doStuff()
    },
    initCallback() {
      SetCallback((data) => {
        this.dcAnswer = data
      })
    }
  }    
}
</script>