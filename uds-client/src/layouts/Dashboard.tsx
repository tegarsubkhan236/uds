import {Outlet, useNavigate} from "react-router-dom";

const dashboard = () => {
    const navigate = useNavigate();

    const handleLogout = () => {
        localStorage.removeItem("token");
        navigate("/login", { replace: true });
    };

    return (
        <div className="container mt-5">
            <h2>Dashboard Layout</h2>
            <button className="btn btn-danger mb-3" onClick={handleLogout}>
                Logout
            </button>
            <Outlet/>
        </div>
    );
};

export default dashboard;
