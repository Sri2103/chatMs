import { BrowserRouter as Router, Routes, Route, } from 'react-router-dom';
import LoginRegister from './pages/LoginRegister';

const Home = () => {
  return (
    <div className="bg-gray-100 h-screen flex justify-center items-center">
      <h1>Welcome to the Home Page</h1>
    </div>
  );
};





const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/authenticate" element={<LoginRegister />} />
      </Routes>
    </Router>
  );
};

export default App;


