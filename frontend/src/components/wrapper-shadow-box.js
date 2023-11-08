import styles from '@/styles/wrapper-shadow-box.module.css'

export default function WrapperShadowBox({ children, alignCenter }) {
  return (
    <div className={alignCenter ? styles.wrap_center : styles.wrapper}>
      {children}
    </div>
  );
}