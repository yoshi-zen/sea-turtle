import styles from '@/styles/layout-12columns.module.css';

export default function Layout12({ children }) {
  return <div className={styles.grid12}>{children}</div>;
}
