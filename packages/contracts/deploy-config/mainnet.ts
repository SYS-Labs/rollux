const config = {
  numDeployConfirmations: 1,
  l1BlockTimeSeconds: 150,
  l2BlockGasLimit: 15_000_000,
  l2ChainId: 570,
  ctcL2GasDiscountDivisor: 32,
  ctcEnqueueGasCost: 60_000,
  sccFaultProofWindowSeconds: 604800,
  sccSequencerPublishWindowSeconds: 12592000,
  ovmSequencerAddress: '0x56Bcd9bf70Cb8c49A7649f223B6100206EC5AC8E',
  ovmProposerAddress: '0x21821Dc24dd29D2D23fE7223dB2BB8D2a72cC95F',
  ovmBlockSignerAddress: '0x00D97b2A26Cb85252998fe7B4bd4eC2118bf6B6E',
  ovmFeeWalletAddress: '0x2fA8986FBf4F9999FBC0CF3f955aDED88444c3EA',
  ovmAddressManagerOwner: '0xbcCC3Ba5e2F84A88d66f62A9fE260A7C303cf440',
  ovmGasPriceOracleOwner: '0xbcCC3Ba5e2F84A88d66f62A9fE260A7C303cf440',
}

export default config
