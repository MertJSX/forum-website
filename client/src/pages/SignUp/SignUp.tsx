import "./SignUp.css";
import { Link } from "react-router";
import { IoArrowBackOutline } from "react-icons/io5";
import { CiLogin } from "react-icons/ci";
import { useState } from "react";
import axios from "axios";

const SignUp = () => {
  const [username, setUsername] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  function trySignUp() {
    axios.post("http://localhost:3000/register-user", {
        username: username,
        email: email,
        password: password
    })
  }

  return (
    <div className="flex mt-28 m-auto justify-center items-center signup-container">
        <div className="bg-gray-800 flex flex-col relative p-2 gap-2 border-r-2 border-blue-300 w-3/6 min-w-[700px] max-w-[1000px] h-[600px] rounded-l-2xl signup-child">
            <div className="absolute z-10 left-25 top-25 select-none">
                <h1 className="text-blue-200 text-5xl z-10 italic font-bold">Become a part of us!</h1>
                <ul  className="text-blue-100 text-2xl list-disc list-inside">
                    <li>Make friends</li>
                    <li>Socialize</li>
                    <li>Ask questions to community</li>
                    <li>Solve your problems</li>
                </ul>
            </div>
            <video 
            src="/videos/signup1.mp4" 
            className="w-full top-0 left-0 absolute h-full object-cover rounded-l-2xl z-0 opacity-45" 
            muted autoPlay loop />
        </div>
        <div className="bg-gray-900 flex flex-col p-2 gap-2 items-center min-w-[400px] w-2/6 max-w-[500px] h-[600px] rounded-r-2xl signup-child">
        <div className="flex">
            <Link 
                to="/"
                className="flex flex-row justify-center items-center min-w-10 text-white text-xl bg-gray-800 hover:bg-gray-700 p-1"
            ><IoArrowBackOutline /></Link>
            <Link 
                to="/signin"
                className="flex flex-row justify-center items-center min-w-24 gap-2 text-white text-xl bg-gray-800 hover:bg-gray-700 p-1"
            ><CiLogin /> Sign in</Link>
        </div>
        <h1 className="text-blue-300 text-5xl font-bold mt-28 italic select-none">SIGN UP</h1>
        <input 
            className="text-center text-lg text-cyan-100 bg-gray-800 w-3/4 rounded-2xl outline-0 focus:bg-gray-700"
            placeholder="Username"
            value={username}
            onChange={
                (e: React.ChangeEvent<HTMLInputElement>) => {
                    setUsername(e.target.value)
            }}
            type="text" />
        <input 
            className="text-center text-lg text-cyan-100 bg-gray-800 w-3/4 rounded-2xl outline-0 focus:bg-gray-700"
            placeholder="Email"
            value={email}
            onChange={
                (e: React.ChangeEvent<HTMLInputElement>) => {
                    setEmail(e.target.value)
            }}
            type="email" />
        <input 
            className="text-center text-lg text-cyan-100 bg-gray-800 w-3/4 rounded-2xl outline-0 focus:bg-gray-700"
            placeholder="Password"
            value={password}
            onChange={
                (e: React.ChangeEvent<HTMLInputElement>) => {
                    setPassword(e.target.value)
            }}
            type="password" />
        <button 
        className="bg-blue-700 hover:bg-blue-600 cursor-pointer text-blue-200 w-3/4 text-lg font-bold rounded-2xl"
        onClick={() => {trySignUp()}}
        type="submit">Sign up</button>

    </div>
    </div>
    
  )
}

export default SignUp