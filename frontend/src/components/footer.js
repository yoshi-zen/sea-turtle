import Container from './container';
import styles from '@/styles/footer.module.css';

export default function Footer() {
  return (
    <footer className={styles.footer}>
      <Container large>
        <div className={styles.flexContainer}>
          <div className={styles.site_icon}></div>
          <ul className={styles.footer_menu}>
            <li>about us</li>
            <li>プライバシーポリシー</li>
            <li>免責事項</li>
            <li>お問い合わせ</li>
          </ul>
          <p>© 2023 ウミガメのスープ</p>
        </div>
      </Container>
    </footer>
  );
}
