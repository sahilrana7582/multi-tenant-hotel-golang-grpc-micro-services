import img1 from "../../assets/login2.jpg"
import img2 from "../../assets/login.jpg"
import img3 from "../../assets/login3.jpg"
import img4 from "../../assets/login4.jpg"
import img5 from "../../assets/login5.jpg"
import LoginForm from '../../features/Auth/components/LoginForm';
import { ErrorFallback } from '../../components/ErrorBoundary/ErrorFallback';
import { ErrorBoundary } from 'react-error-boundary';

const images = [img1, img2, img3, img4, img5];

const Login = () => {

  return (
    <ErrorBoundary
        FallbackComponent={ErrorFallback}
        onReset={() => {
            window.location.reload();
        }}
    >
        <LoginForm  images={images} />
    </ErrorBoundary>
  );
};

export default Login;
