interface AvatarProps {
  src: string;
  alt: string;
  size?: "small" | "medium" | "large";
}

const Avatar: React.FC<AvatarProps> = ({ src, alt, size = "medium" }) => {
  const sizeClasses = {
    small: "w-8 h-8",
    medium: "w-12 h-12",
    large: "w-16 h-16",
  };

  return <img src={src} alt={alt} className={`rounded-full ${sizeClasses[size]}`} />;
};

export default Avatar;
