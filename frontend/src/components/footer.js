import Container from './container';
import styles from '@/styles/footer.module.css';

export default function Footer() {
  return (
    <footer className={styles.footer}>
      <Container large>
        <p>© 2021 ウミガメのスープ</p>
      </Container>
    </footer>
  );
}
