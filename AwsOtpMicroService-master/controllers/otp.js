const mongoose=require("mongoose")
// const contestant = require("../models/contestant")
const otp=require("../models/otp")
var nodemailer = require('nodemailer');

exports.generateOtp=(req,res)=>{
    var otp2=Math.floor(100000 + Math.random() * 900000);
    otp.find({email:req.body.email}).exec((err,otp1)=>{
        if(err){
            error:"Invalid Request"
        }
        if(otp1.length==0){
        const otpData = new otp({
        otp:otp2,
        email:req.body.email
    });
  otpData.save((err, otpData) => {
    if (err) {
      return res.status(400).json({
        error: "NOT able to save otpData in DB"
      });
    }
  });
  res.json("otpData")
    }
    })
}

exports.sendOtp=(req,res)=>{
    //SEND Otp through MAIL
    // console.log("sendOtp",req.body.email);
    var send_otp;
    otp.find({email:req.body.email}).exec((err,userOtp)=>{
        if(err){
            return res.status(400).json({
        error: "Invalid req send otp"
      });
        }
        if(userOtp.length>0){
        //   console.log(userOtp);
        send_otp=userOtp[0].otp;
    var transporter = nodemailer.createTransport({
        service: 'gmail',
        port:465,
        auth: {
          user: 'bigboss5teluguvoting@gmail.com',
          pass: 'pcrsarbmhuyhttoq'
        }
      });
      send();
      async function send() {
        const result = await transporter.sendMail({
            from: 'BigBosstelugu<bigbossvoting@gmail.com>',
            to: req.body.email,
            subject: 'Otp - Quick Byte',
            text: `Your Otp for QuickByte is ${send_otp}`
        });
    
        console.log(JSON.stringify(result, null, 4));
    }
      
    //   transporter.sendMail(mailOptions, function(error, info){
    //       console.log(error);
    //     if (error) {
    //       console.log(error);
    //     } else {
    //       console.log('Email sent: ' + info.response);
    //     }
    //   });
      res.json("otp sent")
    }
    else{
        res.json("otp expired")
    }
  })
}

// exports.sendOtp=(req,res)=>{
//     //SEND Otp through MAIL
//     // console.log("sendOtp",req.body.email);
//     var send_otp;
//     otp.find({email:req.body.email}).exec((err,userOtp)=>{
//         if(err){
//             return res.status(400).json({
//         error: "Invalid req send otp"
//       });
//         }
//         if(userOtp.length>0){
//         //   console.log(userOtp);
//         send_otp=userOtp[0].otp;
//     var transporter = nodemailer.createTransport({
//         service: 'Gmail',
//         port:465,
//         secure:true,
//         auth: {
//           user: 'bigboss5teluguvoting@gmail.com',
//           pass: 'DiNeSh5@'
//         }
//       });
      
//       var mailOptions = {
//         from: 'BigBosstelugu<bigbossvoting@gmail.com>',
//         to: req.body.email,
//         subject: 'Otp for Bigboss 6 telugu',
//         text: `Your Otp is ${send_otp}`
//       };
      
//       transporter.sendMail(mailOptions, function(error, info){
//           console.log(error);
//         if (error) {
//           console.log(error);
//         } else {
//           console.log('Email sent: ' + info.response);
//         }
//       });
//       res.json("otp sent")
//     }
//     else{
//         res.json("otp expired")
//     }
//   })
// }

exports.verifyOtp=(req,res)=>{
  console.log(req.body);
    otp.find({email:req.body.email.email}).exec((err,cate)=>{
        if(err){
            return res.status(400).json({
                error: "INVALID OTP"
            })
        }
        console.log(cate);
        console.log(cate[0]);
        if(cate==null||cate.length==0){
          return res.status(400).json({
                error: "OTP EXPIRED..."
            })
        }
        if(req.body.email.otp==cate[0].otp){
          console.log("sucess");
            return res.json({
                sucess:"Otp Verified"
            })
        }
        else{
            return res.status(400).json({
                error: "INCORRECT OTP"
            })
        }
    })
}
