
FROM registry.innovahq.com/innova/platforms/egonext/tm/common-libs/u22/tm-inetlibs:4.0.00_011 as libs

FROM registry.innovahq.com/innova/platforms/egonext/tm/tm-builder/u22/tm-builder:4.0.00_001 as builder

COPY --from=libs $I_INETLIBS/lib/ $I_INETLIBS/lib
COPY --from=libs $I_INETLIBS/include/ $I_INETLIBS/include
    
