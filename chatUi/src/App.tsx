import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import LoginRegister from "./pages/LoginRegister";
import { ErrorBoundary } from "react-error-boundary";
import { Provider } from "react-redux";
import { store } from "./context/store";
const Home = () => {
  return (
    <div className="bg-gray-100 h-screen flex justify-center items-center">
      <h1>Welcome to the Home Page</h1>
    </div>
  );
};

function Fallback({
  error,
  resetErrorBoundary,
}: {
  error: Error;
  resetErrorBoundary: () => void;
}) {
  // Call resetErrorBoundary() to reset the error boundary and retry the render.

  return (
    <div role="alert">
      <p>Something went wrong:</p>
      <pre style={{ color: "red" }}>{error.message}</pre>
    </div>
  );
}
const App = () => {
  return (
    <ErrorBoundary fallbackRender={Fallback}>
      <Provider store={store}>
        <Router>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/authenticate" element={<LoginRegister />} />
          </Routes>
        </Router>
      </Provider>
    </ErrorBoundary>
  );
};

export default App;
