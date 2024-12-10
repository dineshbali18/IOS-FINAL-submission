const express=require("express");
const router=express.Router();


const {generateOtp,verifyOtp,sendOtp,isPresent}=require('../controllers/otp');

router.post("/user/generateotp",generateOtp)
router.post("/user/sendotp",sendOtp);
router.post("/user/verifyotp",verifyOtp);


module.exports=router;
