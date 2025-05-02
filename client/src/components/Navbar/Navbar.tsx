import GuestNavbarItems from "./GuestNavbarItems";
import UserNavbarItems from "./UserNavbarItems";
import Cookies from "js-cookie";
import "./Navbar.css";

const Navbar = () => {
  return (
    <div className="sticky top-0 left-0 right-0">
      {
        Cookies.get("token") ?
        <UserNavbarItems /> :
        <GuestNavbarItems />
      }
    </div>
  );
};

export default Navbar;
