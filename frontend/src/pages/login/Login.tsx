import styles from './Login.module.scss';
import img1 from "../../assets/login2.jpg"
import img2 from "../../assets/login.jpg"
import img3 from "../../assets/login3.jpg"
import img4 from "../../assets/login4.jpg"
import img5 from "../../assets/login5.jpg"
import { FaGithub, FaGoogle, FaRegEye, FaRegEyeSlash } from 'react-icons/fa';
import { useState } from 'react';
import ImageSlideshow from '../../components/ImageSlideshow/ImageSlideshow';

const images = [img1, img2, img3, img4, img5];

const Login = () => {





    /*---- State Variables ----*/
    const[showPassword, setShowPassword] = useState(false);



    /*---- Functions ----*/
    const toggleShowPassword = () => 
        setShowPassword((prev: boolean): boolean => !prev);
      



  return (
    <div className={styles.container}>

        {/* Left: Image Section */}
        <div className={styles.imageSection}>
            <ImageSlideshow images={images} interval={4000} />
        </div>

        {/* Right: Form Section */}
        <div className={styles.formSide}>
            <form className={styles.form}>
                <h1>Create an account</h1>
                <p>Already have an account? <a href="/login">Log in</a></p>

                <div className={styles.nameGroup}>
                    <input type="text" placeholder="First name" />
                    <input type="text" placeholder="Last name" />
                </div>

                <input type="email" placeholder="Email" autoComplete='username' />
                <div className={styles.passwordWrapper}>
                <input
                    type={showPassword ? "password" : "text"}
                    placeholder="Enter your password"
                    autoComplete="current-password"
                />
                <span onClick={toggleShowPassword}>
                     {showPassword ? < FaRegEye/> : <FaRegEyeSlash/>}
                </span>
                </div>


                <label className={styles.checkbox}>
                    <input type="checkbox" />
                    <span>I agree to the <a href="#">Terms & Conditions</a></span>
                </label>

                <button type="submit">Create account</button>

                <div className={styles.divider}>Or register with</div>

                <div className={styles.socialButtons}>
                    <button className={styles.googleBtn}><FaGoogle /> Google</button>
                    <button className={styles.appleBtn}><FaGithub />Github</button>
                </div>
            </form>
        </div>
    </div>
  );
};

export default Login;
