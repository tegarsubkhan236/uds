import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";
import {useAuth} from "./hooks/useAuth.ts";
import Dashboard from "./layouts/Dashboard.tsx";
import Login from "./pages/Login.tsx";
import VoD from "./pages/VoD.tsx";
import Home from "./pages/Home.tsx";
import Streaming from "./pages/Streaming.tsx";

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
                <Route path="/" element={<PrivateRoute/>}>
                    <Route index element={<Home/>}/>
                    <Route path="/vod" element={<VoD/>}/>
                    <Route path="/streaming" element={<Streaming/>}/>
                </Route>
            </Routes>
        </BrowserRouter>
    )
}

export default App
