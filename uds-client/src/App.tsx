import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";
import {useAuth} from "./hooks/useAuth.ts";
import Dashboard from "./layouts/Dashboard.tsx";
import Login from "./pages/Login.tsx";

const PrivateRoute = () => {
    const { isAuthenticated } = useAuth();
    if (!isAuthenticated) {
        return <Navigate to="/login" replace />;
    } else  {
        return <Dashboard/>
    }
};

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/login" element={<Login/>}/>
                <Route path="/" element={<PrivateRoute/>}/>
            </Routes>
        </BrowserRouter>
    )
}

export default App
