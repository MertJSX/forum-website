import { Link } from "react-router";
import { FaHome } from "react-icons/fa";
import { CiLogin } from "react-icons/ci";
import { TiUserAdd } from "react-icons/ti";

const Navbar = () => {
  return (
    <div className="bg-gray-800 flex flex-row justify-center items-center pb-1 gap-2">
      <h1 className="text-2xl text-blue-300 font-bold italic mr-10 select-none">
        Forum
      </h1>
      <Link
        className="flex justify-center items-center gap-2 text-lg text-white border-b-1 hover:border-sky-300 px-4 pt-2 pb-1"
        to="/"
      >
        <FaHome /> Home
      </Link>
      <Link
        className="flex justify-center items-center gap-2 text-lg text-white border-b-1 hover:border-sky-300 px-4 pt-2 pb-1"
        to="/"
      >
        Explore
      </Link>
      <Link
        className="flex justify-center items-center gap-2 text-lg text-white border-b-1 hover:border-sky-300 px-4 pt-2 pb-1"
        to="/signin"
      >
        <CiLogin /> Sign in
      </Link>
      <Link
        className="flex justify-center items-center gap-2 text-lg text-white border-b-1 hover:border-sky-300 px-4 pt-2 pb-1"
        to="/signup"
      >
        <TiUserAdd /> Sign up
      </Link>
    </div>
  );
};

export default Navbar;
