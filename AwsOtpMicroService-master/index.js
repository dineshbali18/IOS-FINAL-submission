// require('dotenv').config()
const express=require("express")
const mongoose=require("mongoose")
const bodyParser=require("body-parser")
const cors=require("cors")


const app=express();
app.use(bodyParser.json())
app.use(cors());

mongoose
  .connect("mongodb+srv://balidinesh2000:Dinesh5@otpcluster.gvfxo.mongodb.net/?retryWrites=true&w=majority&appName=otpcluster", {
    useNewUrlParser: true,
    useUnifiedTopology: true,
    // useCreateIndex: true
  })
  .then(() => {
    console.log("DB CONNECTED");
  })
  .catch((e)=>{
      console.log(e);
      console.log("DB NOT CONNECTED SUCCESFULLY");
  });

const otpRoutes=require("./routes/otp")
var port=`3000`


app.use("/api",otpRoutes);


app.get("/first",(req,res)=>{
  console.log("route middleware")
  var result = JSON.stringify({status:200, message:"First Node JS Application"})
    res.send(result)
})

app.get("/",(req,res)=>{
  return res.json({msg:"finally worked after a lot of work........."})
})


app.listen(port,()=>{
    console.log(`otpService is running at my own configured port ${port}`)
})
