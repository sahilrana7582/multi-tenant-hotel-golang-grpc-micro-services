import styles from './Login.module.scss';
import loginimg from "../../assets/login2.jpg"

const Login = () => {
  return (
    <div className={styles.container}>

        {/* Left: Image Section */}
        <div className={styles.imageSection}>
            <img src={loginimg}></img>
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

                <input type="email" placeholder="Email" />
                <input type="password" placeholder="Enter your password" />

                <label className={styles.checkbox}>
                    <input type="checkbox" />
                    <span>I agree to the <a href="#">Terms & Conditions</a></span>
                </label>

                <button type="submit">Create account</button>

                <div className={styles.divider}>Or register with</div>
            </form>
        </div>
    </div>
  );
};

export default Login;
