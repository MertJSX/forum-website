import Cookies from "js-cookie"
import { useEffect } from "react"
import { useNavigate } from "react-router"

const Logout = () => {
    const navigate = useNavigate();
    useEffect(() => {
        Cookies.remove("token");
        navigate("/signin")
    }, [])
  return (
    <div>
        <h1>Loading...</h1>
    </div>
  )
}

export default Logout