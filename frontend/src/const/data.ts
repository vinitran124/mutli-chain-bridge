const chain = [
  {
    chainId: '97',
    name: 'BSC Testnet',
    img: '',
    bridgeContractAddress: '0xd15e20F76DC9D2f0671a2E20c303c42d1445c8dd',
  },
  {
    chainId: '5555',
    name: 'CVC Testnet',
    img: '',
    bridgeContractAddress: '0x3ea491a040fa14d8df9457f1a9699b3315d482e3',
  },
];

const coin: { [key: string]: { name: string; address: string }[] } = {
  '97': [
    {
      name: 'VINI',
      address: '0x32D7eEE6513224f459D221BfED0C3af45343CbB7',
    },
    {
      name: 'BNB',
      address: '0x0000000000000000000000000000000000000000',
    },
  ],
  '5': [
    {
      name: 'VINI',
      address: '',
    },
    {
      name: 'ETH',
      address: '0x0000000000000000000000000000000000000000',
    },
  ],
  '5555': [
    {
      name: 'VINI',
      address: '0x596c14ba2b6e759c73895ce13e64e49054181a7f',
    },
    {
      name: 'XSR',
      address: '0x0000000000000000000000000000000000000000',
    },
  ],
};

export const Data = {
  chain,
  coin,
};
