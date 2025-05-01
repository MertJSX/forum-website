import { useEffect } from "react";
import Navbar from "../../components/Navbar/Navbar";
import Cookies from "js-cookie";
import { useNavigate } from "react-router";
import axios from "axios";
import { useState } from "react";
import { FaUserAlt } from "react-icons/fa";

const Profile = () => {
  const navigate = useNavigate();
  let [userData, setUserData] = useState<any>(null);
  useEffect(() => {
    const token = Cookies.get("token");
    if (!token) {
      navigate("/login");
    }
    axios.get("/api/profile", { headers: { Authorization: token } })
    .then((res) => {
      console.log(res.data);
      setUserData(res.data);
      
    }).catch((err) => {
      console.log(err);
      if (err.response.status === 401) {
        navigate("/login");
      }
    });
  }, []);
  return (
    <div>
      <Navbar />
      <div className="flex flex-col items-center">
        {/* <h1 className="text-white">Profile</h1> */}
        {/* <p className="text-white">Welcome to your profile page!</p> */}
        {userData && (
          <div className="text-white mt-10">
            <div className="flex justify-center items-center">
              <FaUserAlt className="text-9xl bg-neutral-600 p-5 rounded-4xl"/>
            </div>
            <h1 className="text-center text-white text-3xl">{userData.username}</h1>
            <h2 className="text-center text-white text-xl">{userData.email}</h2>
          </div>
        )}
      </div>
    </div>
  );
};

export default Profile;

