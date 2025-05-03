import { Link, useNavigate } from "react-router";
import { IoArrowBackOutline } from "react-icons/io5";
import { TiUserAdd } from "react-icons/ti";
import Checkbox from "../../components/minimal-components/Checkbox/Checkbox";
import { useState } from "react";
import axios from "axios";
import isValidEmail from "../../utils/isValidEmail";
import Cookies from "js-cookie";

const SignIn = () => {
  const [emailOrUsername, setEmailOrUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [rememberMe, setRememberMe] = useState<boolean>(false);
  const [errorMsg, setErrorMsg] = useState<string>("")
  const navigate = useNavigate()

  function getToken() {
    console.log(emailOrUsername);
    console.log(password);

    setErrorMsg("")

    axios
      .post("/api/login-user", {
        username: !isValidEmail(emailOrUsername) ? emailOrUsername : "",
        email: isValidEmail(emailOrUsername) ? emailOrUsername : "",
        password: password,
      })
      .then((data) => {
        if (data.data.error) {
          setErrorMsg(data.data.msg)
          setPassword("");
          return
        }
        console.log(data.data.token);
        if (rememberMe) {
          Cookies.set("token", data.data.token, { expires: 7 });
        } else {
          Cookies.set("token", data.data.token);
        }
        navigate("/")
        setEmailOrUsername("");
      })
      .catch((err) => {
        console.log(err);
      });
  }

  return (
    <div className="flex mt-28 m-auto justify-center items-center signup-container">
      <div className="bg-gray-800 flex flex-col relative p-2 gap-2 border-r-2 border-blue-300 w-3/6 min-w-[700px] max-w-[1000px] h-[600px] rounded-l-2xl signup-child">
        <div className="absolute z-10 left-25 top-44 select-none">
          <h1 className="text-blue-200 text-5xl z-10 italic font-bold">
            WELCOME AGAIN!
          </h1>
        </div>
        <video
          src="/videos/signin1.mp4"
          className="w-full top-0 left-0 absolute h-full object-cover rounded-l-2xl z-0 opacity-20"
          muted
          autoPlay
          loop
        />
      </div>
      <div className="bg-gray-800 flex flex-col p-2 gap-2 items-center min-w-[400px] w-2/6 max-w-[500px] h-[600px] rounded-r-2xl signup-child">
        <div className="flex">
          <Link
            to="/"
            className="flex flex-row justify-center items-center min-w-10 text-white text-xl bg-gray-700 hover:bg-gray-600 p-1"
          >
            <IoArrowBackOutline />
          </Link>
          <Link
            to="/signup"
            className="flex flex-row justify-center items-center min-w-24 gap-2 text-white text-xl bg-gray-700 hover:bg-gray-600 p-1"
          >
            <TiUserAdd /> Sign up
          </Link>
        </div>
        <h1 className="text-blue-300 text-5xl font-bold mt-28 italic select-none">
          SIGN IN
        </h1>
        <input
          className="text-center text-lg text-cyan-100 bg-gray-700 w-3/4 rounded-2xl outline-0 focus:bg-gray-600"
          placeholder="Username / email"
          value={emailOrUsername}
          onChange={(e) => {
            setEmailOrUsername(e.target.value);
          }}
          type="text"
          autoFocus
        />
        <input
          className="text-center text-lg text-cyan-100 bg-gray-700 w-3/4 rounded-2xl outline-0 focus:bg-gray-600"
          placeholder="Password"
          value={password}
          onChange={(e) => {
            setPassword(e.target.value);
          }}
          type="password"
        />
        <Checkbox
          label="Remember me"
          checked={rememberMe.toString()}
          onChange={() => {
            setRememberMe(!rememberMe);
          }}
        />
        <button
          className="bg-blue-700 hover:bg-blue-600 cursor-pointer text-blue-200 w-3/4 text-lg font-bold rounded-2xl"
          onClick={() => {
            getToken();
          }}
          type="submit"
        >
          Sign in
        </button>
        <h2 className="text-red-400">{errorMsg ? errorMsg : null}</h2>
      </div>
    </div>
  );
};

export default SignIn;
