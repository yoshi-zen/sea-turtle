import WrapperShadowBox from "@/components/wrapper-shadow-box";

export default function Page() {
  return (
    <>
      <WrapperShadowBox alignCenter >
        <h3>ログイン</h3>
        <p>ユーザー名</p>
        <input type="text" />
        <p>パスワード</p>
        <input type="password" />
      </WrapperShadowBox>
    </>
  );
}
