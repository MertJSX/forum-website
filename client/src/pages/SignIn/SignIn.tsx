import { Link } from "react-router";
import { IoArrowBackOutline } from "react-icons/io5";
import { TiUserAdd } from "react-icons/ti";
import Checkbox from "../../components/minimal-components/Checkbox/Checkbox";
import { useState } from "react";

const SignIn = () => {
  const [rememberMe, setRememberMe] = useState(false);

  return (
    <div className="flex mt-28 m-auto justify-center items-center signup-container">
      <div className="bg-gray-800 flex flex-col relative p-2 gap-2 border-r-2 border-blue-300 w-3/6 min-w-[700px] max-w-[1000px] h-[600px] rounded-l-2xl signup-child">
        <div className="absolute z-10 left-25 top-44 select-none">
          <h1 className="text-blue-200 text-5xl z-10 italic font-bold">
            WELCOME AGAIN!
          </h1>
          {/* <ul  className="text-blue-100 text-2xl list-disc list-inside">
                    <li>Make friends</li>
                    <li>Socialize</li>
                    <li>Ask questions to community</li>
                    <li>Solve your problems</li>
                </ul> */}
        </div>
        <video
          src="/videos/signin1.mp4"
          className="w-full top-0 left-0 absolute h-full object-cover rounded-l-2xl z-0 opacity-20"
          muted
          autoPlay
          loop
        />
      </div>
      <div className="bg-gray-900 flex flex-col p-2 gap-2 items-center min-w-[400px] w-2/6 max-w-[500px] h-[600px] rounded-r-2xl signup-child">
        <div className="flex">
          <Link
            to="/"
            className="flex flex-row justify-center items-center min-w-10 text-white text-xl bg-gray-800 hover:bg-gray-700 p-1"
          >
            <IoArrowBackOutline />
          </Link>
          <Link
            to="/signup"
            className="flex flex-row justify-center items-center min-w-24 gap-2 text-white text-xl bg-gray-800 hover:bg-gray-700 p-1"
          >
            <TiUserAdd /> Sign up
          </Link>
        </div>
        <h1 className="text-blue-300 text-5xl font-bold mt-28 italic select-none">
          SIGN IN
        </h1>
        <input
          className="text-center text-lg text-cyan-100 bg-gray-800 w-3/4 rounded-2xl outline-0 focus:bg-gray-700"
          placeholder="Username / email"
          type="text"
        />
        <input
          className="text-center text-lg text-cyan-100 bg-gray-800 w-3/4 rounded-2xl outline-0 focus:bg-gray-700"
          placeholder="Password"
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
          type="submit"
        >
          Sign in
        </button>
      </div>
    </div>
  );
};

export default SignIn;
