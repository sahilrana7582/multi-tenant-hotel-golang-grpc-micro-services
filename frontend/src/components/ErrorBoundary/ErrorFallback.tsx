import React from "react";
import styles from "./ErrorFallback.module.scss";

type FallbackProps = {
  error: Error;
  resetErrorBoundary: () => void;
};

export const ErrorFallback: React.FC<FallbackProps> = ({ error, resetErrorBoundary }) => {
  return (
    <div className={styles.errorOverlay}>
      <div role="alert" className={styles.errorContainer}>
        <p>Something went wrong:</p>
        <pre>{error.message}</pre>
        <button onClick={resetErrorBoundary}>Try again</button>
      </div>
    </div>
  );
};
