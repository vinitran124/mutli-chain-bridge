export const formatAccount = (walletAddress: string) => {
  const firstFour = walletAddress.substring(0, 6);
  const lastFour = walletAddress.substring(walletAddress.length - 4);
  return `${firstFour}...${lastFour}`;
};
