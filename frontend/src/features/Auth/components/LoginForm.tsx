import { useState } from "react";
import styles from "./LoginForm.module.scss";
import ImageSlideshow from "../../../components/ImageSlideshow/ImageSlideshow";
import { FaRegEye, FaRegEyeSlash } from "react-icons/fa";
import { useForm } from "react-hook-form";
import { loginSchema, type LoginFormProps, type LoginFormValues } from "../type";
import { zodResolver } from "@hookform/resolvers/zod";
import { useLogin } from "../hooks/useLogin";

const LoginForm = ({ images }: LoginFormProps) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormValues>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });


  // Hooks
  
  const {handleLogin, loading} = useLogin();
  const [showPassword, setShowPassword] = useState(false);

  const toggleShowPassword = () => setShowPassword((prev) => !prev);



  return (
    <div className={styles.container}>
      <div className={styles.imageSection}>
        <ImageSlideshow images={images} interval={4000} />
      </div>

      <div className={styles.formSide}>
        <form className={styles.form} onSubmit={handleSubmit(handleLogin)}>
          <h1 className={styles.heading}>Login</h1>
          <p className={styles.subtext}>Look inside what’s happening!</p>


          <div className={styles.field}>
            <input
              {...register("email")}
              type="email"
              placeholder="Email"
              className={styles.input}
            />
            {errors.email && (
              <p className={styles.errorMsg}>{errors.email.message}</p>
            )}
          </div>

          <div className={styles.field}>
            <div className={styles.passwordWrapper}>
              <input
                {...register("password")}
                type={showPassword ? "text" : "password"}
                placeholder="Enter your password"
                className={styles.input}
              />
              <span className={styles.icon} onClick={toggleShowPassword}>
                {showPassword ? <FaRegEyeSlash /> : <FaRegEye />}
              </span>
            </div>
            {errors.password && (
              <p className={styles.errorMsg}>{errors.password.message}</p>
            )}
          </div>

          <button
            type="submit"
            className={styles.button}
            disabled={loading}
          >
            {loading ? "Logging in..." : "Login"}
          </button>
        </form>
      </div>
    </div>
  );
};

export default LoginForm;
