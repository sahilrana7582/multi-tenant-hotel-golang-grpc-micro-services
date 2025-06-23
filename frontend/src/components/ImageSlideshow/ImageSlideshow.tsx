import React, { useEffect, useState } from "react";
import styles from "./ImageSlideshow.module.scss";

interface ImageSlideshowProps {
  images: string[];
  interval?: number;
}

const ImageSlideshow: React.FC<ImageSlideshowProps> = ({
  images,
  interval = 4000,
}) => {
  const [currentIndex, setCurrentIndex] = useState(0);

  useEffect(() => {
    const timer = setInterval(() => {
      setCurrentIndex((prev) => (prev + 1) % images.length);
    }, interval);

    return () => clearInterval(timer);
  }, [images.length, interval]);

  return (
    <div className={styles.slideshow}>
      {images.map((img, index) => (
        <img
          key={index}
          src={img}
          alt={`Slide ${index + 1}`}
          className={`${styles.image} ${index === currentIndex ? styles.active : ""}`}
        />
      ))}
    </div>
  );
};

export default ImageSlideshow;
