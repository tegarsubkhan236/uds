import {NavLink, Outlet, useNavigate} from "react-router-dom";
import {useAuthApi} from "../hooks/useAuthApi.ts";

const Dashboard = () => {
    const navigate = useNavigate();

    const { logoutHandler, auth } = useAuthApi();

    const handleLogout = async () => {
        await logoutHandler();
        navigate("/login", {replace: true});
    };

    return (
        <div className="container-fluid">
            <nav className="navbar navbar-expand-lg bg-body-tertiary">
                <div className="container-fluid">
                    <NavLink className="navbar-brand" to="/">UDS</NavLink>
                    <button className="navbar-toggler" type="button" data-bs-toggle="collapse"
                            data-bs-target="#navbarNav"
                            aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    <div className="collapse navbar-collapse" id="navbarNav">
                        <ul className="navbar-nav">
                            <li className="nav-item">
                                <NavLink to="/" className={({ isActive }) => `nav-link${isActive ? ' active' : ''}`}>
                                    Home
                                </NavLink>
                            </li>
                            <li className="nav-item">
                                <NavLink to="/vod" className={({ isActive }) => `nav-link${isActive ? ' active' : ''}`}>
                                    VoD
                                </NavLink>
                            </li>
                            <li className="nav-item">
                                <NavLink to="/streaming" className={({ isActive }) => `nav-link${isActive ? ' active' : ''}`}>
                                    Streaming
                                </NavLink>
                            </li>
                        </ul>
                        <div className="ms-auto d-flex align-items-center gap-2">
                            <span className="navbar-text fw-bold">
                                Hallo {auth?.username}
                            </span>
                            <button className="btn btn-outline-danger" onClick={handleLogout}>
                                Logout
                            </button>
                        </div>
                    </div>
                </div>
            </nav>
            <Outlet/>
        </div>
    );
};

export default Dashboard;
